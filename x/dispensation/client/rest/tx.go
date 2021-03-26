package rest

// The packages below are commented out at first to prevent an error if this file isn't initially saved.
import (
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc(
		"/dispensation/airdrop",
		airdropHandler(cliCtx),
	).Methods("POST")
}

type (
	AirdropReq struct {
		BaseReq rest.BaseReq `json:"base_req"`
		Signer  string       `json:"signer"` // User who is trying to create the pool // ExternalAsset in the pool pair (ex rwn:ceth)
		Input   string       `json:"input"`  // NativeAssetAmount is the amount of native asset being added
		Output  string       `json:"output"` // ExternalAssetAmount is the amount of external asset being added
	}
)

func airdropHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var req CreatePoolReq
		//if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, "failed to parse request")
		//	return
		//}
		//baseReq := req.BaseReq.Sanitize()
		//if !baseReq.ValidateBasic(w) {
		//	return
		//}
		//signer, err := sdk.AccAddressFromBech32(req.Signer)
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}
		//msg := types.NewMsgCreatePool(signer, req.ExternalAsset, req.NativeAssetAmount, req.ExternalAssetAmount)
		//err = msg.ValidateBasic()
		//if err != nil {
		//	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		//	return
		//}
		//utils.WriteGenerateStdTxResponse(w, cliCtx, baseReq, []sdk.Msg{msg})
	}
}
