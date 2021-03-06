package types

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	hub "github.com/sentinel-official/hub/types"
)

func TestMsgRegisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgRegisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgRegisterNode(nil, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgRegisterNode([]byte(""), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"node_type is empty",
			NewMsgRegisterNode(TestAddress1, "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("type"),
		}, {
			"version is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("version"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", nil, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is negative",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthNeg, "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"internet_speed is zero",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthZero, "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, ""),
			ErrorInvalidField("encryption"),
		}, {
			"valid",
			NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgRegisterNode_GetSignBytes(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgRegisterNode_GetSigners(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgRegisterNode_Type(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, "register_node", msg.Type())
}

func TestMsgRegisterNode_Route(t *testing.T) {
	msg := NewMsgRegisterNode(TestAddress1, "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgUpdateNodeInfo_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgUpdateNodeInfo
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgUpdateNodeInfo(nil, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgUpdateNodeInfo([]byte(""), hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("from"),
		}, {
			"node_moniker length is greater than 128",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", strings.Repeat("X", 130), sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("moniker"),
		}, {
			"prices_per_gb is nil",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", nil, TestBandwidthPos1, "encryption"),
			nil,
		}, {
			"prices_per_gb is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is negative",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.Coin{"stake", sdk.NewInt(-100)}}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"prices_per_gb is zero",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 0)}, TestBandwidthPos1, "encryption"),
			ErrorInvalidField("prices_per_gb"),
		}, {
			"internet_speed is zero",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthZero, "encryption"),
			nil,
		}, {
			"internet_speed is negative",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthNeg, "encryption"),
			ErrorInvalidField("internet_speed"),
		}, {
			"encryption is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, ""),
			nil,
		}, {
			"type is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			nil,
		}, {
			"version is empty",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			nil,
		}, {
			"valid",
			NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption"),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgUpdateNode_GetSignBytes(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgUpdateNode_GetSigners(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgUpdateNode_Type(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, "update_node_info", msg.Type())
}

func TestMsgUpdateNode_Route(t *testing.T) {
	msg := NewMsgUpdateNodeInfo(TestAddress1, hub.NewNodeID(1), "node_type", "version", "moniker", sdk.Coins{sdk.NewInt64Coin("stake", 100)}, TestBandwidthPos1, "encryption")
	require.Equal(t, RouterKey, msg.Route())
}

func TestMsgDeregisterNode_ValidateBasic(t *testing.T) {
	tests := []struct {
		name string
		msg  *MsgDeregisterNode
		want sdk.Error
	}{
		{
			"from is nil",
			NewMsgDeregisterNode(nil, hub.NewNodeID(1)),
			ErrorInvalidField("from"),
		}, {
			"from is empty",
			NewMsgDeregisterNode([]byte(""), hub.NewNodeID(1)),
			ErrorInvalidField("from"),
		}, {
			"valid",
			NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1)),
			nil,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.msg.ValidateBasic(); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("\ngot = %vwant = %v", got, tc.want)
			}
		})
	}
}

func TestMsgDeregisterNode_GetSignBytes(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	msgBytes, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}

	require.Equal(t, msgBytes, msg.GetSignBytes())
}

func TestMsgDeregisterNode_GetSigners(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	require.Equal(t, []sdk.AccAddress{TestAddress1}, msg.GetSigners())
}

func TestMsgDeregisterNode_Type(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	require.Equal(t, "deregister_node", msg.Type())
}

func TestMsgDeregisterNode_Route(t *testing.T) {
	msg := NewMsgDeregisterNode(TestAddress1, hub.NewNodeID(1))
	require.Equal(t, RouterKey, msg.Route())
}
