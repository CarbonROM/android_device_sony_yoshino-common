#
# Copyright (C) 2011 The Android Open-Source Project
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

LOCAL_PATH := $(call my-dir)

include $(call all-makefiles-under,$(LOCAL_PATH))

include $(CLEAR_VARS)
# Create mount points for oem configs
OEM_MOUNT_POINTS := $(TARGET_ROOT_OUT)/oem

ALL_DEFAULT_INSTALLED_MODULES += $(OEM_MOUNT_POINTS)

$(OEM_MOUNT_POINTS):
	@echo "Creating $(OEM_MOUNT_POINTS) subdirs"
	@mkdir -p $(TARGET_ROOT_OUT)/oem/modem-config
	@mkdir -p $(TARGET_ROOT_OUT)/oem/system-properties
