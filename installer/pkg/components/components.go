package components

import (
	"github.com/gitpod-io/observability/installer/pkg/common"
	nodeExporter "github.com/gitpod-io/observability/installer/pkg/components/node-exporter"
	"github.com/gitpod-io/observability/installer/pkg/components/prometheusOperator"
	"github.com/gitpod-io/observability/installer/pkg/components/pyrra"
)

var MonitoringCentralObjects = common.MergeLists(pyrra.Objects)
var MonitoringSatelliteObjects = common.MergeLists(pyrra.Objects, nodeExporter.Objects, prometheusOperator.Objects)
