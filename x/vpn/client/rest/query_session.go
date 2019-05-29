package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	sdkTypes "github.com/ironman0x7b2/sentinel-sdk/types"
	"github.com/ironman0x7b2/sentinel-sdk/x/vpn/client/common"
)

func getSessionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := sdkTypes.NewIDFromString(vars["sessionID"])
		session, err := common.QuerySession(cliCtx, cdc, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, session, cliCtx.Indent)
	}
}

func getSessionsOfSubscriptionHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		id := sdkTypes.NewIDFromString(vars["subscriptionID"])
		sessions, err := common.QuerySessionsOfSubscription(cliCtx, cdc, id)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, sessions, cliCtx.Indent)
	}
}

func getAllSessionsHandlerFunc(cliCtx context.CLIContext, cdc *codec.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sessions, err := common.QueryAllSessions(cliCtx, cdc)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponse(w, cdc, sessions, cliCtx.Indent)
	}
}
