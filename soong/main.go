package yoshino

import (
    "android/soong/android"
)

func init() {
    android.RegisterModuleType("sony_yoshino_init_library", initLibraryFactory)
}
