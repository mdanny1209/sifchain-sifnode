/**
 * This script will pause the bridgebank smart contract to shut down EVM imports.
 * Use with care.
 * 
 * To use this script make sure the environment variable ACTIVE_PRIVATE_KEY is set to the pauser private key before running
 * Make sure you have environment variable MAINNET_URL set to a HTTP/HTTPS Full/Archive Node for the EVM chain in communication
 * Make sure the pauser address has at least 1 ETH before running
 */

import { ethers } from "hardhat";
import { BridgeBank__factory } from "../build";

const bridgeBankAddress = process.env["BRIDGEBANK_ADDRESS"] ?? "0xB5F54ac4466f5ce7E0d8A5cB9FE7b8c0F35B7Ba8";

async function pauseBridge() {
    // Connect to BridgeBank and get User Info
    const bridgebankFactory = await ethers.getContractFactory("BridgeBank") as BridgeBank__factory;
    const bridgebank = await bridgebankFactory.attach(bridgeBankAddress);
    const userAddress = await bridgebank.signer.getAddress();
    let paused = await bridgebank.paused();
    const balance = await bridgebank.signer.getBalance();

    // Sanity Condition Checks
    if (paused) {
        console.error("Bridgebank is already paused, no actions to do");
        return;
    }

    if (!(await bridgebank.pausers(userAddress))) {
        console.error(`Private key has public address: ${userAddress}, which is not a valid pauser address`);
        return;
    }

    if (balance.lt(ethers.utils.parseEther("0.5"))) {
        console.error(`Script requires a minimum of 0.5 ETH before it will attempt to run.`)
        console.error(`Current balance is ${ethers.utils.formatEther(balance)} ETH`);
        return;
    }

    // Pause The Bridge
    console.log("Sending the pause transaction");
    const tx = await bridgebank.pause();
    console.log("Transaction sent, waiting for transaction receipt");
    const receipt = await tx.wait();
    console.log(`Received transaction receipt. Transaction Hash: ${receipt.transactionHash}`);

    // Confirm the bridge is now paused
    paused = await bridgebank.paused();
    if (paused) {
        console.log("Confirmed the BridgeBank is now paused");
    } else {
        console.error("We have received a pause transaction receipt but BridgeBank is not paused...");
        console.error("!!!!!CRITICAL CONDITION REACHED, CONTACT PEGGY TEAM IMMEDIATELY!!!!!")
    }
}

pauseBridge()
  .then(() => console.log("Pauser Script has completed"))
  .catch((error) => console.error("An error ocurred when attempting to pause bridgebank: ", error))