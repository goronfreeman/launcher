//go:build windows
// +build windows

package table

import (
	"github.com/kolide/launcher/pkg/osquery/tables/dsim_default_associations"
	"github.com/kolide/launcher/pkg/osquery/tables/secedit"
	"github.com/kolide/launcher/pkg/osquery/tables/wifi_networks"
	"github.com/kolide/launcher/pkg/osquery/tables/windowsupdatetable"
	"github.com/kolide/launcher/pkg/osquery/tables/wmitable"

	"github.com/go-kit/kit/log"
	osquery "github.com/osquery/osquery-go"
)

func platformTables(client *osquery.ExtensionManagerClient, logger log.Logger, currentOsquerydBinaryPath string) []osquery.OsqueryPlugin {
	return []osquery.OsqueryPlugin{
		ProgramIcons(),
		dsim_default_associations.TablePlugin(client, logger),
		secedit.TablePlugin(client, logger),
		wifi_networks.TablePlugin(client, logger),
		windowsupdatetable.TablePlugin(windowsupdatetable.UpdatesTable, client, logger),
		windowsupdatetable.TablePlugin(windowsupdatetable.HistoryTable, client, logger),
		wmitable.TablePlugin(client, logger),
	}
}
