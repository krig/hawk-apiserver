package api

import "strings"

// api/v1/configuration/cluster_property
func handleConfigCluster(urllist []string, cib *Cib) (bool, interface{}) {
	return true, FetchNV(cib.Configuration.CrmConfig.ClusterPropertySet)
}

// api/v1/configuration/rsc_defaults
func handleConfigRscDefaults(urllist []string, cib *Cib) (bool, interface{}) {
	return true, FetchNV(cib.Configuration.RscDefaults)
}

// api/v1/configuration/op_defaults
func handleConfigOpDefaults(urllist []string, cib *Cib) (bool, interface{}) {
	return true, FetchNV(cib.Configuration.OpDefaults)
}

// api/v1/status/summary
func handleStateSummary(urllist []string, crmMon *CrmMon) (bool, interface{}) {
	summaryData := crmMon.CrmMonSummary
	if summaryData == nil {
		return true, nil
	}

	ch := make(chan string)
	go FetchContent(ch, GetNumField(summaryData), summaryData)
	nv := make(map[string]string)
	for n := range ch {
		res := strings.Split(n, ";")
		// in crm_mon xml outputs, there are some elements which have some attributes
		// for the reason about namespace, I combine the element tag and each attribute name
		// so the output looks flat, hope this can be useful
		key := res[2] + "_" + res[0]
		nv[key] = res[1]
	}

	return true, nv
}
