var ua = navigator.userAgent;
var isAndroid = /Android/i.test(ua);
var isBlackBerry = /BlackBerry/i.test(ua);
var isWindowPhone = /IEMobile/i.test(ua);
var isIOS = /iPhone|iPad|iPod/i.test(ua);
isMobile = isAndroid || isBlackBerry || isWindowPhone || isIOS;

var jd_union_pid = "948016030";
var jd_union_euid = "";

if (isMobile) {
    // 小米手环2抢券9块9
    var jd_union_pid = "944528180";
    var jd_union_euid = "";
}