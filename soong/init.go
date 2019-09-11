package yoshino

import (
    "android/soong/android"
    "android/soong/cc"
    "strings"
)

func initLibs(ctx android.BaseContext) []string {
    var libs []string

    var config = ctx.AConfig().VendorConfig("SONY_YOSHINO_INIT")
    var extension = strings.TrimSpace(config.String("EXTENSION"))

    if len(extension) > 0 {
        libs = append(libs, extension)
    }
    return libs
}

func initLibrary(ctx android.LoadHookContext) {
    type props struct {
        Target struct {
            Android struct {
                Whole_static_libs []string
            }
        }
    }

    p := &props{}
    p.Target.Android.Whole_static_libs = initLibs(ctx)
    ctx.AppendProperties(p)
}

func initLibraryFactory() android.Module {
    module, library := cc.NewLibrary(android.HostAndDeviceSupported)
    library.BuildOnlyStatic()
    newMod := module.Init()
    android.AddLoadHook(newMod, initLibrary)
    return newMod
}
