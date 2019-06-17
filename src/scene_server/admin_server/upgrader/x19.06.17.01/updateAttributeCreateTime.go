/*
 * Tencent is pleased to support the open source community by making 蓝鲸 available.
 * Copyright (C) 2017-2018 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 */

package x19_06_17_01

import (
	"configcenter/src/common"
	"configcenter/src/scene_server/admin_server/upgrader"
	"configcenter/src/storage/dal"
	"context"
)

func updateInnerModelPreField(ctx context.Context, db dal.RDB, conf *upgrader.Config) error {
	filter := map[string]interface{}{
		common.BKObjIDField: map[string]interface{}{
			common.BKDBIN: []string{
				common.BKInnerObjIDSwitch,
				common.BKInnerObjIDHost,
				common.BKInnerObjIDRouter,
				common.BKInnerObjIDBlance,
				common.BKInnerObjIDFirewall,
				common.BKInnerObjIDWeblogic,
				common.BKInnerObjIDApache,
				common.BKInnerObjIDTomcat,
			},
		},
	}
	attribute := map[string]interface{}{
		common.BKIsPre: true,
	}
	err := db.Table(common.BKTableNameObjDes).Update(ctx, filter, &attribute)
	if err != nil {
		return err
	}

	return nil

}
