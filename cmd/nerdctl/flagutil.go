/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package main

import (
	"github.com/containerd/nerdctl/pkg/api/types"
	"github.com/spf13/cobra"
)

func processRootCmdFlags(cmd *cobra.Command) (types.GlobalCommandOptions, error) {
	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	debugFull, err := cmd.Flags().GetBool("debug-full")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	address, err := cmd.Flags().GetString("address")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	namespace, err := cmd.Flags().GetString("namespace")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	snapshotter, err := cmd.Flags().GetString("snapshotter")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	cniPath, err := cmd.Flags().GetString("cni-path")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	cniConfigPath, err := cmd.Flags().GetString("cni-netconfpath")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	dataRoot, err := cmd.Flags().GetString("data-root")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	cgroupManager, err := cmd.Flags().GetString("cgroup-manager")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	insecureRegistry, err := cmd.Flags().GetBool("insecure-registry")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	hostsDir, err := cmd.Flags().GetStringSlice("hosts-dir")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	experimental, err := cmd.Flags().GetBool("experimental")
	if err != nil {
		return types.GlobalCommandOptions{}, err
	}
	return types.GlobalCommandOptions{
		Debug:            debug,
		DebugFull:        debugFull,
		Address:          address,
		Namespace:        namespace,
		Snapshotter:      snapshotter,
		CNIPath:          cniPath,
		CNINetConfPath:   cniConfigPath,
		DataRoot:         dataRoot,
		CgroupManager:    cgroupManager,
		InsecureRegistry: insecureRegistry,
		HostsDir:         hostsDir,
		Experimental:     experimental,
	}, nil
}
