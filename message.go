package mipush

import (
	"net/url"
	"strconv"
	"time"
)

type PassThroughType int8

var (
	PassThroughNotify    PassThroughType = 0 // 通知栏消息
	PassThroughPenetrate PassThroughType = 1 // 透传消息
)

type NotifyTypeType int8

var (
	NotifyTypeAll     NotifyTypeType = -1
	NotifyTypeSound   NotifyTypeType = 1 // 使用默认提示音提示
	NotifyTypeVibrate NotifyTypeType = 2 // 使用默认震动提示
	NotifyTypeLights  NotifyTypeType = 4 // 使用默认led灯光提示
)

type NotifyEffectType string

var (
	NotifyLauncherActive NotifyEffectType = "1" // 打开当前app对应的Launcher Activity
	NotifyActivity       NotifyEffectType = "2" // 打开当前app内的任意一个Activity
	NotifyWeb            NotifyEffectType = "3" // 打开网页
)

type Message url.Values

func NewMessage() Message {
	return make(Message)
}

// 消息的内容。
func (this Message) Payload(payload string) Message {
	message := url.Values(this)
	message.Set("payload", payload)
	return this
}

// pass_through的值可以为：
// 0 表示通知栏消息
// 1 表示透传消息
func (this Message) PassThrough(passThroughID PassThroughType) Message {
	message := url.Values(this)
	message.Set("pass_through", strconv.FormatInt(int64(passThroughID), 10))
	return this
}

// 通知栏展示的通知的标题。
func (this Message) Title(title string) Message {
	message := url.Values(this)
	message.Set("title", title)
	return this
}

// 通知栏展示的通知的描述。
func (this Message) Description(description string) Message {
	message := url.Values(this)
	message.Set("description", description)
	return this
}

// notify_type的值可以是DEFAULT_ALL或者以下其他几种的OR组合：
// DEFAULT_ALL = -1;
// DEFAULT_SOUND  = 1;  // 使用默认提示音提示；
// DEFAULT_VIBRATE = 2;  // 使用默认震动提示；
// DEFAULT_LIGHTS = 4;   // 使用默认led灯光提示；
func (this Message) NotifyType(notifyType NotifyTypeType) Message {
	message := url.Values(this)
	message.Set("notify_type", strconv.FormatInt(int64(notifyType), 10))
	return this
}

// 可选项。如果用户离线，设置消息在服务器保存的时间，单位：ms。服务器默认最长保留两周。
func (this Message) TimeToLive(timeToLive time.Duration) Message {
	// 低于1毫秒不进行赋值
	if timeToLive < time.Millisecond {
		return this
	}
	timeToLive /= time.Millisecond
	message := url.Values(this)
	message.Set("time_to_live", strconv.FormatInt(int64(timeToLive), 10))
	return this
}

// 可选项。定时发送消息。用自1970年1月1日以来00:00:00.0 UTC时间表示（以毫秒为单位的时间）。注：仅支持七天内的定时消息。
func (this Message) TimeToSend(timeToSend int64) Message {
	message := url.Values(this)
	message.Set("time_to_send", strconv.FormatInt(timeToSend, 10))
	return this

}

// 可选项。默认情况下，通知栏只显示一条推送消息。如果通知栏要显示多条推送消息，需要针对不同的消息设置不同的notify_id（相同notify_id的通知栏消息会覆盖之前的）。
func (this Message) NotifyID(notifyID uint64) Message {
	message := url.Values(this)
	message.Set("notify_id", strconv.FormatUint(notifyID, 10))
	return this
}

// 可选项，自定义通知栏消息铃声。extra.sound_uri的值设置为铃声的URI。参考2.2.1注：铃声文件放在Android app的raw目录下。
func (this Message) ExtraSoundURI(soundURI string) Message {
	message := url.Values(this)
	message.Set("extra.sound_uri", soundURI)
	return this
}

// 可选项，开启通知消息在状态栏滚动显示。
func (this Message) ExtraTicker(ticker string) Message {
	message := url.Values(this)
	message.Set("extra.ticker", ticker)
	return this
}

// 可选项，开启/关闭app在前台时的通知弹出。当extra.notify_foreground值为”1″时，app会弹出通知栏消息；当extra.notify_foreground值为”0″时，app不会弹出通知栏消息。注：默认情况下会弹出通知栏消息。参考2.2.2
func (this Message) ExtraNotifyForeground(notifyForeground string) Message {
	message := url.Values(this)
	message.Set("extra.notify_foreground", notifyForeground)
	return this
}

// 可选项 打开App 的 Launcher Activity
func (this Message) OpenLauncherActivity() Message {
	message := url.Values(this)
	message.Set("extra.notify_effect", string(NotifyLauncherActive))
	return this
}

// 可选项 打开App 内 某个 Activity
func (this Message) OpenActivity(intentURI string) Message {
	message := url.Values(this)
	message.Set("extra.intent_uri", intentURI)
	message.Set("extra.notify_effect", string(NotifyActivity))
	return this
}

// 可选项 打开网页
func (this Message) OpenWebURI(webURI string) Message {
	message := url.Values(this)
	message.Set("extra.web_uri", webURI)
	message.Set("extra.notify_effect", string(NotifyWeb))
	return this
}

// 	可选项，控制是否需要进行平缓发送。默认不支持。value代表平滑推送的速度。注：服务端支持最低1000/s的qps，最高100000/s。也就是说，如果将平滑推送设置为500，那么真实的推送速度是3000/s，如果大于1000小于100000，则以设置的qps为准。
func (this Message) ExtraFlowControl(flowControl int) Message {
	message := url.Values(this)
	message.Set("extra.flow_control", strconv.FormatInt(int64(flowControl), 10))
	return this
}

// 可选项，自定义通知栏样式，设置为客户端要展示的layout文件名。
func (this Message) ExtraLayoutName(layoutName string) Message {
	message := url.Values(this)
	message.Set("extra.layout_name", layoutName)
	return this
}

// 可选项，自定义通知栏样式，必须与layout_name一同使用，指定layout中各控件展示的内容。
func (this Message) ExtraLayoutValue(layoutValue string) Message {
	message := url.Values(this)
	message.Set("extra.layout_value", layoutValue)
	return this
}

// 	可选项，使用推送批次（JobKey）功能聚合消息。客户端会按照jobkey去重，即相同jobkey的消息只展示第一条，其他的消息会被忽略。合法的jobkey由数字（[0-9]），大小写字母（[a-zA-Z]），下划线（_）和中划线（-）组成，长度不大于8个字符。
func (this Message) ExtraJobKey(jobKey string) Message {
	message := url.Values(this)
	message.Set("extra.jobkey", jobKey)
	return this
}

// 可选项，开启消息送达和点击回执。将extra.callback的值设置为第三方接收回执的http接口
func (this Message) ExtraCallback(callback string) Message {
	message := url.Values(this)
	message.Set("extra.callback", callback)
	return this
}

// 可选项，可以接收消息的设备的语言范围，用逗号分隔
func (this Message) ExtraLocale(locale string) Message {
	message := url.Values(this)
	message.Set("extra.locale", locale)
	return this
}

// 可选项，无法收到消息的设备的语言范围，逗号分隔
func (this Message) ExtraLocaleNotIn(localeNotIn string) Message {
	message := url.Values(this)
	message.Set("extra.locale_not_in", localeNotIn)
	return this
}

// 可选项，model支持三种用法。
// 可以收到消息的设备的机型范围，逗号分隔。当前设备的model的获取方法：
// Build.MODEL
// 比如，小米手机4移动版用”MI 4LTE”表示。
// 可以收到消息的设备的品牌范围，逗号分割。比如，三星手机用”samsung”表示。 （目前覆盖42个主流品牌，对应关系见附录”品牌表”）
// 可以收到消息的设备的价格范围，逗号分隔。比如，0-999价格的设备用”0-999″表示。 （目前覆盖4个价格区间，对应关系见附录”价格表”）
func (this Message) ExtraModel(model string) Message {
	message := url.Values(this)
	message.Set("extra.model", model)
	return this
}

// 可选项，无法收到消息的设备的机型范围，逗号分隔
func (this Message) ExtraModelNotIn(modelNotIn string) Message {
	message := url.Values(this)
	message.Set("extra.model_not_in", modelNotIn)
	return this
}

// 可以接收消息的app版本号，用逗号分割。安卓app版本号来源于manifest文件中的”android:versionName”的值。注：目前支持MiPush_SDK_Client_2_2_12_sdk.jar（及以后）的版本。
func (this Message) ExtraAppVersion(appVersion string) Message {
	message := url.Values(this)
	message.Set("extra.app_version", appVersion)
	return this
}

// 无法接收消息的app版本号，用逗号分割。
func (this Message) ExtraAppVersionNotIn(appVersionNotIn string) Message {
	message := url.Values(this)
	message.Set("extra.app_version_not_in", appVersionNotIn)
	return this
}

// 可选项，指定在特定的网络环境下才能接收到消息。目前仅支持指定”wifi”。
func (this Message) ExtraConnpt(connpt string) Message {
	message := url.Values(this)
	message.Set("extra.​connpt", connpt)
	return this
}
