package yoshino

import (
    "android/soong/android"
    "android/soong/cc"
    "strconv"
    "strings"
)

func initSrc(ctx android.BaseContext) []string {
    var src []string

    var config = ctx.AConfig().VendorConfig("SONY_YOSHINO_INIT")
    var extension = strings.TrimSpace(config.String("EXTENSION"))

    src = append(src, extension)
    return src
}

func initBinary(ctx android.LoadHookContext) {
    type props struct {
        Target struct {
            Android struct {
                Src []string
            }
        }
    }

    p := &props{}
    p.Target.Android.Src = initSrc(ctx)
    ctx.AppendProperties(p)
}

func initBinaryFactory() android.Module {
    module, _ := cc.NewBinary(android.HostAndDeviceSupported)
    newMod := module.Init()
    android.AddLoadHook(newMod, initBinary)
    return newMod
}
