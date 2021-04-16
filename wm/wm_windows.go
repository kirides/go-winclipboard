package wm

import "strconv"

// Contains some Window Messages extracted from WIN 10 SDK

// Aswell as the WM-type that implements fmt.Stringer
// which lets you do wm.WM(1).String() = "CREATE"

// Window Messages (WM_)
const (
	NULL                    = 0x0000
	CREATE                  = 0x0001
	DESTROY                 = 0x0002
	MOVE                    = 0x0003
	SIZE                    = 0x0005
	ACTIVATE                = 0x0006
	SETFOCUS                = 0x0007
	KILLFOCUS               = 0x0008
	ENABLE                  = 0x000A
	SETREDRAW               = 0x000B
	SETTEXT                 = 0x000C
	GETTEXT                 = 0x000D
	GETTEXTLENGTH           = 0x000E
	PAINT                   = 0x000F
	CLOSE                   = 0x0010
	QUERYENDSESSION         = 0x0011
	QUERYOPEN               = 0x0013
	ENDSESSION              = 0x0016
	QUIT                    = 0x0012
	ERASEBKGND              = 0x0014
	SYSCOLORCHANGE          = 0x0015
	SHOWWINDOW              = 0x0018
	WININICHANGE            = 0x001A
	SETTINGCHANGE           = WININICHANGE
	DEVMODECHANGE           = 0x001B
	ACTIVATEAPP             = 0x001C
	FONTCHANGE              = 0x001D
	TIMECHANGE              = 0x001E
	CANCELMODE              = 0x001F
	SETCURSOR               = 0x0020
	MOUSEACTIVATE           = 0x0021
	CHILDACTIVATE           = 0x0022
	QUEUESYNC               = 0x0023
	GETMINMAXINFO           = 0x0024
	PAINTICON               = 0x0026
	ICONERASEBKGND          = 0x0027
	NEXTDLGCTL              = 0x0028
	SPOOLERSTATUS           = 0x002A
	DRAWITEM                = 0x002B
	MEASUREITEM             = 0x002C
	DELETEITEM              = 0x002D
	VKEYTOITEM              = 0x002E
	CHARTOITEM              = 0x002F
	SETFONT                 = 0x0030
	GETFONT                 = 0x0031
	SETHOTKEY               = 0x0032
	GETHOTKEY               = 0x0033
	QUERYDRAGICON           = 0x0037
	COMPAREITEM             = 0x0039
	GETOBJECT               = 0x003D
	COMPACTING              = 0x0041
	COMMNOTIFY              = 0x0044
	WINDOWPOSCHANGING       = 0x0046
	WINDOWPOSCHANGED        = 0x0047
	POWER                   = 0x0048
	COPYDATA                = 0x004A
	CANCELJOURNAL           = 0x004B
	NOTIFY                  = 0x004E
	INPUTLANGCHANGEREQUEST  = 0x0050
	INPUTLANGCHANGE         = 0x0051
	TCARD                   = 0x0052
	HELP                    = 0x0053
	USERCHANGED             = 0x0054
	NOTIFYFORMAT            = 0x0055
	CONTEXTMENU             = 0x007B
	STYLECHANGING           = 0x007C
	STYLECHANGED            = 0x007D
	DISPLAYCHANGE           = 0x007E
	GETICON                 = 0x007F
	SETICON                 = 0x0080
	NCCREATE                = 0x0081
	NCDESTROY               = 0x0082
	NCCALCSIZE              = 0x0083
	NCHITTEST               = 0x0084
	NCPAINT                 = 0x0085
	NCACTIVATE              = 0x0086
	GETDLGCODE              = 0x0087
	SYNCPAINT               = 0x0088
	NCMOUSEMOVE             = 0x00A0
	NCLBUTTONDOWN           = 0x00A1
	NCLBUTTONUP             = 0x00A2
	NCLBUTTONDBLCLK         = 0x00A3
	NCRBUTTONDOWN           = 0x00A4
	NCRBUTTONUP             = 0x00A5
	NCRBUTTONDBLCLK         = 0x00A6
	NCMBUTTONDOWN           = 0x00A7
	NCMBUTTONUP             = 0x00A8
	NCMBUTTONDBLCLK         = 0x00A9
	KEYFIRST                = 0x0100
	KEYDOWN                 = 0x0100
	KEYUP                   = 0x0101
	CHAR                    = 0x0102
	DEADCHAR                = 0x0103
	SYSKEYDOWN              = 0x0104
	SYSKEYUP                = 0x0105
	SYSCHAR                 = 0x0106
	SYSDEADCHAR             = 0x0107
	IME_STARTCOMPOSITION    = 0x010D
	IME_ENDCOMPOSITION      = 0x010E
	IME_COMPOSITION         = 0x010F
	IME_KEYLAST             = 0x010F
	INITDIALOG              = 0x0110
	COMMAND                 = 0x0111
	SYSCOMMAND              = 0x0112
	TIMER                   = 0x0113
	HSCROLL                 = 0x0114
	VSCROLL                 = 0x0115
	INITMENU                = 0x0116
	INITMENUPOPUP           = 0x0117
	GESTURE                 = 0x0119
	GESTURENOTIFY           = 0x011A
	MENUSELECT              = 0x011F
	MENUCHAR                = 0x0120
	ENTERIDLE               = 0x0121
	MENURBUTTONUP           = 0x0122
	MENUDRAG                = 0x0123
	MENUGETOBJECT           = 0x0124
	UNINITMENUPOPUP         = 0x0125
	MENUCOMMAND             = 0x0126
	CTLCOLORMSGBOX          = 0x0132
	CTLCOLOREDIT            = 0x0133
	CTLCOLORLISTBOX         = 0x0134
	CTLCOLORBTN             = 0x0135
	CTLCOLORDLG             = 0x0136
	CTLCOLORSCROLLBAR       = 0x0137
	CTLCOLORSTATIC          = 0x0138
	MOUSEFIRST              = 0x0200
	MOUSEMOVE               = 0x0200
	LBUTTONDOWN             = 0x0201
	LBUTTONUP               = 0x0202
	LBUTTONDBLCLK           = 0x0203
	RBUTTONDOWN             = 0x0204
	RBUTTONUP               = 0x0205
	RBUTTONDBLCLK           = 0x0206
	MBUTTONDOWN             = 0x0207
	MBUTTONUP               = 0x0208
	MBUTTONDBLCLK           = 0x0209
	PARENTNOTIFY            = 0x0210
	ENTERMENULOOP           = 0x0211
	EXITMENULOOP            = 0x0212
	NEXTMENU                = 0x0213
	SIZING                  = 0x0214
	CAPTURECHANGED          = 0x0215
	MOVING                  = 0x0216
	POWERBROADCAST          = 0x0218
	DEVICECHANGE            = 0x0219
	MDICREATE               = 0x0220
	MDIDESTROY              = 0x0221
	MDIACTIVATE             = 0x0222
	MDIRESTORE              = 0x0223
	MDINEXT                 = 0x0224
	MDIMAXIMIZE             = 0x0225
	MDITILE                 = 0x0226
	MDICASCADE              = 0x0227
	MDIICONARRANGE          = 0x0228
	MDIGETACTIVE            = 0x0229
	MDISETMENU              = 0x0230
	ENTERSIZEMOVE           = 0x0231
	EXITSIZEMOVE            = 0x0232
	DROPFILES               = 0x0233
	MDIREFRESHMENU          = 0x0234
	POINTERDEVICECHANGE     = 0x238
	POINTERDEVICEINRANGE    = 0x239
	POINTERDEVICEOUTOFRANGE = 0x23A
	TOUCH                   = 0x0240
	NCPOINTERUPDATE         = 0x0241
	NCPOINTERDOWN           = 0x0242
	NCPOINTERUP             = 0x0243
	POINTERUPDATE           = 0x0245
	POINTERDOWN             = 0x0246
	POINTERUP               = 0x0247
	POINTERENTER            = 0x0249
	POINTERLEAVE            = 0x024A
	POINTERACTIVATE         = 0x024B
	POINTERCAPTURECHANGED   = 0x024C
	TOUCHHITTESTING         = 0x024D
	POINTERWHEEL            = 0x024E
	POINTERHWHEEL           = 0x024F
	POINTERROUTEDTO         = 0x0251
	POINTERROUTEDAWAY       = 0x0252
	POINTERROUTEDRELEASED   = 0x0253
	IME_SETCONTEXT          = 0x0281
	IME_NOTIFY              = 0x0282
	IME_CONTROL             = 0x0283
	IME_COMPOSITIONFULL     = 0x0284
	IME_SELECT              = 0x0285
	IME_CHAR                = 0x0286
	IME_REQUEST             = 0x0288
	IME_KEYDOWN             = 0x0290
	IME_KEYUP               = 0x0291
	NCMOUSEHOVER            = 0x02A0
	NCMOUSELEAVE            = 0x02A2
	DPICHANGED              = 0x02E0
	DPICHANGED_BEFOREPARENT = 0x02E2
	DPICHANGED_AFTERPARENT  = 0x02E3
	GETDPISCALEDSIZE        = 0x02E4
	CUT                     = 0x0300
	COPY                    = 0x0301
	PASTE                   = 0x0302
	CLEAR                   = 0x0303
	UNDO                    = 0x0304
	RENDERFORMAT            = 0x0305
	RENDERALLFORMATS        = 0x0306
	DESTROYCLIPBOARD        = 0x0307
	DRAWCLIPBOARD           = 0x0308
	PAINTCLIPBOARD          = 0x0309
	VSCROLLCLIPBOARD        = 0x030A
	SIZECLIPBOARD           = 0x030B
	ASKCBFORMATNAME         = 0x030C
	CHANGECBCHAIN           = 0x030D
	HSCROLLCLIPBOARD        = 0x030E
	QUERYNEWPALETTE         = 0x030F
	PALETTEISCHANGING       = 0x0310
	PALETTECHANGED          = 0x0311
	HOTKEY                  = 0x0312
	PRINT                   = 0x0317
	PRINTCLIENT             = 0x0318
	DWMNCRENDERINGCHANGED   = 0x031F
	GETTITLEBARINFOEX       = 0x033F
	HANDHELDFIRST           = 0x0358
	HANDHELDLAST            = 0x035F
	AFXFIRST                = 0x0360
	AFXLAST                 = 0x037F
	PENWINFIRST             = 0x0380
	PENWINLAST              = 0x038F
	APP                     = 0x8000
	USER                    = 0x0400
	CLIPBOARDUPDATE         = 0x031D
)

var mapWMMsg = map[uint32]string{
	NULL:                    "WM_NULL",
	CREATE:                  "WM_CREATE",
	DESTROY:                 "WM_DESTROY",
	MOVE:                    "WM_MOVE",
	SIZE:                    "WM_SIZE",
	ACTIVATE:                "WM_ACTIVATE",
	SETFOCUS:                "WM_SETFOCUS",
	KILLFOCUS:               "WM_KILLFOCUS",
	ENABLE:                  "WM_ENABLE",
	SETREDRAW:               "WM_SETREDRAW",
	SETTEXT:                 "WM_SETTEXT",
	GETTEXT:                 "WM_GETTEXT",
	GETTEXTLENGTH:           "WM_GETTEXTLENGTH",
	PAINT:                   "WM_PAINT",
	CLOSE:                   "WM_CLOSE",
	QUERYENDSESSION:         "WM_QUERYENDSESSION",
	QUERYOPEN:               "WM_QUERYOPEN",
	ENDSESSION:              "WM_ENDSESSION",
	QUIT:                    "WM_QUIT",
	ERASEBKGND:              "WM_ERASEBKGND",
	SYSCOLORCHANGE:          "WM_SYSCOLORCHANGE",
	SHOWWINDOW:              "WM_SHOWWINDOW",
	WININICHANGE:            "WM_WININICHANGE",
	DEVMODECHANGE:           "WM_DEVMODECHANGE",
	ACTIVATEAPP:             "WM_ACTIVATEAPP",
	FONTCHANGE:              "WM_FONTCHANGE",
	TIMECHANGE:              "WM_TIMECHANGE",
	CANCELMODE:              "WM_CANCELMODE",
	SETCURSOR:               "WM_SETCURSOR",
	MOUSEACTIVATE:           "WM_MOUSEACTIVATE",
	CHILDACTIVATE:           "WM_CHILDACTIVATE",
	QUEUESYNC:               "WM_QUEUESYNC",
	GETMINMAXINFO:           "WM_GETMINMAXINFO",
	PAINTICON:               "WM_PAINTICON",
	ICONERASEBKGND:          "WM_ICONERASEBKGND",
	NEXTDLGCTL:              "WM_NEXTDLGCTL",
	SPOOLERSTATUS:           "WM_SPOOLERSTATUS",
	DRAWITEM:                "WM_DRAWITEM",
	MEASUREITEM:             "WM_MEASUREITEM",
	DELETEITEM:              "WM_DELETEITEM",
	VKEYTOITEM:              "WM_VKEYTOITEM",
	CHARTOITEM:              "WM_CHARTOITEM",
	SETFONT:                 "WM_SETFONT",
	GETFONT:                 "WM_GETFONT",
	SETHOTKEY:               "WM_SETHOTKEY",
	GETHOTKEY:               "WM_GETHOTKEY",
	QUERYDRAGICON:           "WM_QUERYDRAGICON",
	COMPAREITEM:             "WM_COMPAREITEM",
	GETOBJECT:               "WM_GETOBJECT",
	COMPACTING:              "WM_COMPACTING",
	COMMNOTIFY:              "WM_COMMNOTIFY",
	WINDOWPOSCHANGING:       "WM_WINDOWPOSCHANGING",
	WINDOWPOSCHANGED:        "WM_WINDOWPOSCHANGED",
	POWER:                   "WM_POWER",
	COPYDATA:                "WM_COPYDATA",
	CANCELJOURNAL:           "WM_CANCELJOURNAL",
	NOTIFY:                  "WM_NOTIFY",
	INPUTLANGCHANGEREQUEST:  "WM_INPUTLANGCHANGEREQUEST",
	INPUTLANGCHANGE:         "WM_INPUTLANGCHANGE",
	TCARD:                   "WM_TCARD",
	HELP:                    "WM_HELP",
	USERCHANGED:             "WM_USERCHANGED",
	NOTIFYFORMAT:            "WM_NOTIFYFORMAT",
	CONTEXTMENU:             "WM_CONTEXTMENU",
	STYLECHANGING:           "WM_STYLECHANGING",
	STYLECHANGED:            "WM_STYLECHANGED",
	DISPLAYCHANGE:           "WM_DISPLAYCHANGE",
	GETICON:                 "WM_GETICON",
	SETICON:                 "WM_SETICON",
	NCCREATE:                "WM_NCCREATE",
	NCDESTROY:               "WM_NCDESTROY",
	NCCALCSIZE:              "WM_NCCALCSIZE",
	NCHITTEST:               "WM_NCHITTEST",
	NCPAINT:                 "WM_NCPAINT",
	NCACTIVATE:              "WM_NCACTIVATE",
	GETDLGCODE:              "WM_GETDLGCODE",
	SYNCPAINT:               "WM_SYNCPAINT",
	NCMOUSEMOVE:             "WM_NCMOUSEMOVE",
	NCLBUTTONDOWN:           "WM_NCLBUTTONDOWN",
	NCLBUTTONUP:             "WM_NCLBUTTONUP",
	NCLBUTTONDBLCLK:         "WM_NCLBUTTONDBLCLK",
	NCRBUTTONDOWN:           "WM_NCRBUTTONDOWN",
	NCRBUTTONUP:             "WM_NCRBUTTONUP",
	NCRBUTTONDBLCLK:         "WM_NCRBUTTONDBLCLK",
	NCMBUTTONDOWN:           "WM_NCMBUTTONDOWN",
	NCMBUTTONUP:             "WM_NCMBUTTONUP",
	NCMBUTTONDBLCLK:         "WM_NCMBUTTONDBLCLK",
	KEYDOWN:                 "WM_KEYDOWN",
	KEYUP:                   "WM_KEYUP",
	CHAR:                    "WM_CHAR",
	DEADCHAR:                "WM_DEADCHAR",
	SYSKEYDOWN:              "WM_SYSKEYDOWN",
	SYSKEYUP:                "WM_SYSKEYUP",
	SYSCHAR:                 "WM_SYSCHAR",
	SYSDEADCHAR:             "WM_SYSDEADCHAR",
	IME_STARTCOMPOSITION:    "WM_IME_STARTCOMPOSITION",
	IME_ENDCOMPOSITION:      "WM_IME_ENDCOMPOSITION",
	IME_COMPOSITION:         "WM_IME_COMPOSITION",
	INITDIALOG:              "WM_INITDIALOG",
	COMMAND:                 "WM_COMMAND",
	SYSCOMMAND:              "WM_SYSCOMMAND",
	TIMER:                   "WM_TIMER",
	HSCROLL:                 "WM_HSCROLL",
	VSCROLL:                 "WM_VSCROLL",
	INITMENU:                "WM_INITMENU",
	INITMENUPOPUP:           "WM_INITMENUPOPUP",
	GESTURE:                 "WM_GESTURE",
	GESTURENOTIFY:           "WM_GESTURENOTIFY",
	MENUSELECT:              "WM_MENUSELECT",
	MENUCHAR:                "WM_MENUCHAR",
	ENTERIDLE:               "WM_ENTERIDLE",
	MENURBUTTONUP:           "WM_MENURBUTTONUP",
	MENUDRAG:                "WM_MENUDRAG",
	MENUGETOBJECT:           "WM_MENUGETOBJECT",
	UNINITMENUPOPUP:         "WM_UNINITMENUPOPUP",
	MENUCOMMAND:             "WM_MENUCOMMAND",
	CTLCOLORMSGBOX:          "WM_CTLCOLORMSGBOX",
	CTLCOLOREDIT:            "WM_CTLCOLOREDIT",
	CTLCOLORLISTBOX:         "WM_CTLCOLORLISTBOX",
	CTLCOLORBTN:             "WM_CTLCOLORBTN",
	CTLCOLORDLG:             "WM_CTLCOLORDLG",
	CTLCOLORSCROLLBAR:       "WM_CTLCOLORSCROLLBAR",
	CTLCOLORSTATIC:          "WM_CTLCOLORSTATIC",
	MOUSEMOVE:               "WM_MOUSEMOVE",
	LBUTTONDOWN:             "WM_LBUTTONDOWN",
	LBUTTONUP:               "WM_LBUTTONUP",
	LBUTTONDBLCLK:           "WM_LBUTTONDBLCLK",
	RBUTTONDOWN:             "WM_RBUTTONDOWN",
	RBUTTONUP:               "WM_RBUTTONUP",
	RBUTTONDBLCLK:           "WM_RBUTTONDBLCLK",
	MBUTTONDOWN:             "WM_MBUTTONDOWN",
	MBUTTONUP:               "WM_MBUTTONUP",
	MBUTTONDBLCLK:           "WM_MBUTTONDBLCLK",
	PARENTNOTIFY:            "WM_PARENTNOTIFY",
	ENTERMENULOOP:           "WM_ENTERMENULOOP",
	EXITMENULOOP:            "WM_EXITMENULOOP",
	NEXTMENU:                "WM_NEXTMENU",
	SIZING:                  "WM_SIZING",
	CAPTURECHANGED:          "WM_CAPTURECHANGED",
	MOVING:                  "WM_MOVING",
	POWERBROADCAST:          "WM_POWERBROADCAST",
	DEVICECHANGE:            "WM_DEVICECHANGE",
	MDICREATE:               "WM_MDICREATE",
	MDIDESTROY:              "WM_MDIDESTROY",
	MDIACTIVATE:             "WM_MDIACTIVATE",
	MDIRESTORE:              "WM_MDIRESTORE",
	MDINEXT:                 "WM_MDINEXT",
	MDIMAXIMIZE:             "WM_MDIMAXIMIZE",
	MDITILE:                 "WM_MDITILE",
	MDICASCADE:              "WM_MDICASCADE",
	MDIICONARRANGE:          "WM_MDIICONARRANGE",
	MDIGETACTIVE:            "WM_MDIGETACTIVE",
	MDISETMENU:              "WM_MDISETMENU",
	ENTERSIZEMOVE:           "WM_ENTERSIZEMOVE",
	EXITSIZEMOVE:            "WM_EXITSIZEMOVE",
	DROPFILES:               "WM_DROPFILES",
	MDIREFRESHMENU:          "WM_MDIREFRESHMENU",
	POINTERDEVICECHANGE:     "WM_POINTERDEVICECHANGE",
	POINTERDEVICEINRANGE:    "WM_POINTERDEVICEINRANGE",
	POINTERDEVICEOUTOFRANGE: "WM_POINTERDEVICEOUTOFRANGE",
	TOUCH:                   "WM_TOUCH",
	NCPOINTERUPDATE:         "WM_NCPOINTERUPDATE",
	NCPOINTERDOWN:           "WM_NCPOINTERDOWN",
	NCPOINTERUP:             "WM_NCPOINTERUP",
	POINTERUPDATE:           "WM_POINTERUPDATE",
	POINTERDOWN:             "WM_POINTERDOWN",
	POINTERUP:               "WM_POINTERUP",
	POINTERENTER:            "WM_POINTERENTER",
	POINTERLEAVE:            "WM_POINTERLEAVE",
	POINTERACTIVATE:         "WM_POINTERACTIVATE",
	POINTERCAPTURECHANGED:   "WM_POINTERCAPTURECHANGED",
	TOUCHHITTESTING:         "WM_TOUCHHITTESTING",
	POINTERWHEEL:            "WM_POINTERWHEEL",
	POINTERHWHEEL:           "WM_POINTERHWHEEL",
	POINTERROUTEDTO:         "WM_POINTERROUTEDTO",
	POINTERROUTEDAWAY:       "WM_POINTERROUTEDAWAY",
	POINTERROUTEDRELEASED:   "WM_POINTERROUTEDRELEASED",
	IME_SETCONTEXT:          "WM_IME_SETCONTEXT",
	IME_NOTIFY:              "WM_IME_NOTIFY",
	IME_CONTROL:             "WM_IME_CONTROL",
	IME_COMPOSITIONFULL:     "WM_IME_COMPOSITIONFULL",
	IME_SELECT:              "WM_IME_SELECT",
	IME_CHAR:                "WM_IME_CHAR",
	IME_REQUEST:             "WM_IME_REQUEST",
	IME_KEYDOWN:             "WM_IME_KEYDOWN",
	IME_KEYUP:               "WM_IME_KEYUP",
	NCMOUSEHOVER:            "WM_NCMOUSEHOVER",
	NCMOUSELEAVE:            "WM_NCMOUSELEAVE",
	DPICHANGED:              "WM_DPICHANGED",
	DPICHANGED_BEFOREPARENT: "WM_DPICHANGED_BEFOREPARENT",
	DPICHANGED_AFTERPARENT:  "WM_DPICHANGED_AFTERPARENT",
	GETDPISCALEDSIZE:        "WM_GETDPISCALEDSIZE",
	CUT:                     "WM_CUT",
	COPY:                    "WM_COPY",
	PASTE:                   "WM_PASTE",
	CLEAR:                   "WM_CLEAR",
	UNDO:                    "WM_UNDO",
	RENDERFORMAT:            "WM_RENDERFORMAT",
	RENDERALLFORMATS:        "WM_RENDERALLFORMATS",
	DESTROYCLIPBOARD:        "WM_DESTROYCLIPBOARD",
	DRAWCLIPBOARD:           "WM_DRAWCLIPBOARD",
	PAINTCLIPBOARD:          "WM_PAINTCLIPBOARD",
	VSCROLLCLIPBOARD:        "WM_VSCROLLCLIPBOARD",
	SIZECLIPBOARD:           "WM_SIZECLIPBOARD",
	ASKCBFORMATNAME:         "WM_ASKCBFORMATNAME",
	CHANGECBCHAIN:           "WM_CHANGECBCHAIN",
	HSCROLLCLIPBOARD:        "WM_HSCROLLCLIPBOARD",
	QUERYNEWPALETTE:         "WM_QUERYNEWPALETTE",
	PALETTEISCHANGING:       "WM_PALETTEISCHANGING",
	PALETTECHANGED:          "WM_PALETTECHANGED",
	HOTKEY:                  "WM_HOTKEY",
	PRINT:                   "WM_PRINT",
	PRINTCLIENT:             "WM_PRINTCLIENT",
	GETTITLEBARINFOEX:       "WM_GETTITLEBARINFOEX",
	HANDHELDFIRST:           "WM_HANDHELDFIRST",
	HANDHELDLAST:            "WM_HANDHELDLAST",
	AFXFIRST:                "WM_AFXFIRST",
	AFXLAST:                 "WM_AFXLAST",
	PENWINFIRST:             "WM_PENWINFIRST",
	PENWINLAST:              "WM_PENWINLAST",
	APP:                     "WM_APP",
	USER:                    "WM_USER",
	CLIPBOARDUPDATE:         "WM_CLIPBOARDUPDATE",
	DWMNCRENDERINGCHANGED:   "WM_DWMNCRENDERINGCHANGED",
}

type WM uint32

func (wm WM) String() string {
	if name, ok := mapWMMsg[uint32(wm)]; ok {
		return name
	} else {
		return strconv.FormatUint(uint64(wm), 10)
	}
}
func String(wm uint32) string {
	return WM(wm).String()
}

