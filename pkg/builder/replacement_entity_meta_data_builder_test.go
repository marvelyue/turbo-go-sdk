package builder

import (
	"reflect"
	"testing"

	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
	"github.com/turbonomic/turbo-go-sdk/util/rand"
)

func TestNewReplacementEntityMetaDataBuilder(t *testing.T) {
	expectedEntityMetaDataBuilder := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp:  []string{},
			BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	meta := NewReplacementEntityMetaDataBuilder()
	if !reflect.DeepEqual(meta, expectedEntityMetaDataBuilder) {
		t.Errorf("Expected %+v, got %+v", expectedEntityMetaDataBuilder, meta)
	}
}

func TestReplacementEntityMetaDataBuilder_Build(t *testing.T) {
	base := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp:  []string{},
			BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	expected := &proto.EntityDTO_ReplacementEntityMetaData{
		IdentifyingProp:  []string{},
		BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
	}
	metadata := base.Build()
	if !reflect.DeepEqual(expected, metadata) {
		t.Errorf("Expected %+v, got %+v", expected, metadata)
	}
}

func TestReplacementEntityMetaDataBuilder_Matching(t *testing.T) {
	base := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp:  []string{},
			BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	property := rand.String(5)
	expectedBuilder := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp:  []string{property},
			BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	builder := base.Matching(property)
	if !reflect.DeepEqual(builder, expectedBuilder) {
		t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
	}
}

func TestReplacementEntityMetaDataBuilder_PatchBuying(t *testing.T) {
	base := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp:  []string{},
			BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	commType := rand.RandomCommodityType()
	expectedBuilder := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp: []string{},
			BuyingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{
				{
					CommodityType: &commType,
				},
			},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	builder := base.PatchBuying(commType)
	if !reflect.DeepEqual(builder, expectedBuilder) {
		t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
	}
}

func TestReplacementEntityMetaDataBuilder_PatchSelling(t *testing.T) {
	base := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp:  []string{},
			BuyingCommTypes:  []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
		},
	}
	commType := rand.RandomCommodityType()
	expectedBuilder := &ReplacementEntityMetaDataBuilder{
		&proto.EntityDTO_ReplacementEntityMetaData{
			IdentifyingProp: []string{},
			BuyingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{},
			SellingCommTypes: []*proto.EntityDTO_ReplacementCommodityPropertyData{
				{
					CommodityType: &commType,
				},
			},
		},
	}
	builder := base.PatchSelling(commType)
	if !reflect.DeepEqual(builder, expectedBuilder) {
		t.Errorf("Expected %+v, got %+v", expectedBuilder, builder)
	}
}
