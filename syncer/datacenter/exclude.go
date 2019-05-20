/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package datacenter

import (
	scpb "github.com/apache/servicecomb-service-center/server/core/proto"
	pb "github.com/apache/servicecomb-service-center/syncer/proto"
)

// exclude find out services data belonging to self
func (s *store) exclude(data *pb.SyncData, curMapping pb.SyncMapping) {
	services := make([]*pb.SyncService, 0, 10)
	for _, svc := range data.Services {
		svc.Instances = s.excludeInstances(svc.Instances, curMapping)
		services = append(services, svc)
	}
	data.Services = services
}

// excludeInstances find out the instance data belonging to self, through the mapping table
func (s *store) excludeInstances(ins []*scpb.MicroServiceInstance, curMapping pb.SyncMapping) []*scpb.MicroServiceInstance {
	nis := make([]*scpb.MicroServiceInstance, 0, len(ins))
	for _, inst := range ins {
		if index := curMapping.CurrentIndex(inst.InstanceId); index != -1 {
			continue
		}
		nis = append(nis, inst)
	}
	return nis
}
