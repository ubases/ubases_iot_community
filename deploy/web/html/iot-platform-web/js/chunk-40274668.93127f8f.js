(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-40274668","chunk-f274e0ee"],{"4c97":function(t,e,a){},"706e":function(t,e,a){"use strict";a("4c97")},a467:function(t,e,a){"use strict";a.r(e);a("b0c0");var r=function(){var t=this,e=t._self._c;return e("a-modal",{attrs:{title:"导出设备",width:800,visible:t.visible,"confirm-loading":t.confirmLoading},on:{cancel:t.handleCancel},scopedSlots:t._u([{key:"footer",fn:function(){return[e("a-button",{key:"back",on:{click:t.handleCancel}},[t._v("取消")]),e("a-divider",{attrs:{type:"vertical"}}),e("a-button",{key:"submit",attrs:{type:"primary",loading:t.confirmLoading},on:{click:t.handleOk}},[t._v(" 确定 ")])]},proxy:!0}])},[e("a-spin",{attrs:{spinning:t.confirmLoading}},[e("a-row",{attrs:{gutter:16}},[e("a-col",{attrs:{md:8,sm:24}},[e("a-form-item",[e("a-select",{staticStyle:{width:"100%"},attrs:{placeholder:"是否绑定"},model:{value:t.queryParam.isActive,callback:function(e){t.$set(t.queryParam,"isActive",e)},expression:"queryParam.isActive"}},[e("a-select-option",{attrs:{value:1}},[t._v("已激活")]),e("a-select-option",{attrs:{value:2}},[t._v("未激活")])],1)],1)],1),e("a-col",{attrs:{md:8,sm:24}},[e("a-form-item",[e("a-select",{staticStyle:{width:"100%"},attrs:{placeholder:"是否在线"},model:{value:t.queryParam.isOnline,callback:function(e){t.$set(t.queryParam,"isOnline",e)},expression:"queryParam.isOnline"}},[e("a-select-option",{attrs:{value:1}},[t._v("在线")]),e("a-select-option",{attrs:{value:2}},[t._v("离线")])],1)],1)],1),e("a-col",{attrs:{md:8,sm:24}},[e("a-form-item",[e("a-select",{staticStyle:{width:"100%"},attrs:{placeholder:"设备归属"},model:{value:t.queryParam.deviceNature,callback:function(e){t.$set(t.queryParam,"deviceNature",e)},expression:"queryParam.deviceNature"}},[e("a-select-option",{attrs:{value:1}},[t._v("开发者")]),e("a-select-option",{attrs:{value:2}},[t._v("公司名称&账号")]),e("a-select-option",{attrs:{value:3}},[t._v("个人账号")])],1)],1)],1),e("a-col",{attrs:{md:8,sm:24}},[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入设备ID，设备名称"},model:{value:t.queryParam.searchKey,callback:function(e){t.$set(t.queryParam,"searchKey",e)},expression:"queryParam.searchKey"}})],1)],1),e("a-col",{attrs:{md:8,sm:24}},[e("a-form-item",[e("a-select",{staticStyle:{width:"100%"},attrs:{placeholder:"所属产品"},model:{value:t.queryParam.productId,callback:function(e){t.$set(t.queryParam,"productId",e)},expression:"queryParam.productId"}},t._l(t.dataSource,(function(a){return e("a-select-option",{key:a.id,attrs:{value:a.id}},[t._v(" "+t._s(a.name)+" ")])})),1)],1)],1),e("a-col",{attrs:{md:19,sm:24}},[e("a-form-item",[e("a-date-picker",{staticStyle:{width:"48%"},attrs:{format:"YYYY-MM-DD",placeholder:"请选择开始时间","disabled-date":t.disabledStartDate},on:{openChange:t.handleStartOpenChange},model:{value:t.startValue,callback:function(e){t.startValue=e},expression:"startValue"}}),e("span",{staticStyle:{width:"4%",display:"inline-block","text-align":"center"}},[t._v(" ~ ")]),e("a-date-picker",{staticStyle:{width:"48%"},attrs:{format:"YYYY-MM-DD",placeholder:"请选择结束时间","disabled-date":t.disabledEndDate,open:t.endOpen},on:{openChange:t.handleEndOpenChange},model:{value:t.endValue,callback:function(e){t.endValue=e},expression:"endValue"}})],1)],1),e("a-col",{attrs:{md:8,sm:24}},[e("a-form-item",[e("a-button",{attrs:{icon:"search",type:"primary"},on:{click:t.queryList}},[t._v(" 查询 ")]),e("a-divider",{attrs:{type:"vertical"}}),e("a-button",{attrs:{icon:"redo"},on:{click:t.reset}},[t._v(" 重置 ")])],1)],1)],1),e("div",{staticClass:"log-total"},[t._v(" 导出设备 共查询到"+t._s(t.logTotal)+"条数据，请确认是否全部导出 ")])],1)],1)},n=[],o=a("5530"),i=(a("d3b7"),a("c4c8")),u=(a("c1df"),{data:function(){return{visible:!1,confirmLoading:!1,endValue:"",startValue:"",queryParam:{},endOpen:!1,logTotal:0,dataSource:[]}},watch:{visible:function(t,e){t&&(this.queryParam={},this.logTotal=0)}},created:function(){this.init()},methods:{init:function(){var t=this;Object(i["I"])({}).then((function(e){0==e.code?t.dataSource=e.data.list||[]:t.dataSource=[]}))},queryList:function(){var t=this;this.queryParam.endTime=Date.parse(this.endValue)/1e3,this.queryParam.startTime=Date.parse(this.startValue)/1e3,this.confirmLoading=!0,Object(i["z"])({query:this.queryParam,isOnlyCount:1,searchKey:this.queryParam.searchKey}).then((function(e){0==e.code&&(t.logTotal=e.data.total)})).finally((function(){t.confirmLoading=!1}))},reset:function(){this.queryParam={},this.endValue="",this.startValue="",this.queryList()},handleOk:function(){var t=this;this.logTotal?(this.confirmLoading=!0,this.$DownloadTemplate(this,{url:"/v1/platform/web/iot/activeDevice/platformExport",method:"post",query:Object(o["a"])({},this.queryParam),searchKey:this.queryParam.searchKey},"设备列表.xlsx").finally((function(){t.confirmLoading=!1}))):this.$message.error("请先查询设备")},handleCancel:function(){this.visible=!1},disabledStartDate:function(t){var e=this.endValue;return!(!t||!e)&&t.valueOf()>e.valueOf()},disabledEndDate:function(t){var e=this.startValue;return!(!t||!e)&&e.valueOf()>=t.valueOf()},handleStartOpenChange:function(t){t||(this.endOpen=!0)},handleEndOpenChange:function(t){this.endOpen=t}}}),c=u,l=(a("706e"),a("2877")),d=Object(l["a"])(c,r,n,!1,null,"3e3952f7",null);e["default"]=d.exports},c4c8:function(t,e,a){"use strict";a.d(e,"u",(function(){return n})),a.d(e,"t",(function(){return o})),a.d(e,"b",(function(){return i})),a.d(e,"l",(function(){return u})),a.d(e,"h",(function(){return c})),a.d(e,"D",(function(){return l})),a.d(e,"M",(function(){return d})),a.d(e,"e",(function(){return s})),a.d(e,"o",(function(){return m})),a.d(e,"E",(function(){return p})),a.d(e,"O",(function(){return f})),a.d(e,"P",(function(){return v})),a.d(e,"d",(function(){return y})),a.d(e,"n",(function(){return b})),a.d(e,"q",(function(){return h})),a.d(e,"I",(function(){return q})),a.d(e,"G",(function(){return w})),a.d(e,"H",(function(){return k})),a.d(e,"N",(function(){return g})),a.d(e,"j",(function(){return O})),a.d(e,"f",(function(){return _})),a.d(e,"p",(function(){return P})),a.d(e,"s",(function(){return x})),a.d(e,"g",(function(){return j})),a.d(e,"r",(function(){return S})),a.d(e,"a",(function(){return C})),a.d(e,"k",(function(){return D})),a.d(e,"K",(function(){return T})),a.d(e,"z",(function(){return L})),a.d(e,"x",(function(){return $})),a.d(e,"A",(function(){return N})),a.d(e,"B",(function(){return I})),a.d(e,"C",(function(){return V})),a.d(e,"y",(function(){return K})),a.d(e,"F",(function(){return E})),a.d(e,"w",(function(){return z})),a.d(e,"i",(function(){return F})),a.d(e,"c",(function(){return R})),a.d(e,"m",(function(){return Y})),a.d(e,"v",(function(){return A})),a.d(e,"L",(function(){return M})),a.d(e,"J",(function(){return G}));var r=a("b775");function n(t){return Object(r["b"])({url:"/v1/platform/web/pm/productType/get",method:"post",data:t})}function o(t){return Object(r["b"])({url:"/v1/platform/web/pm/productType/get",method:"get",params:t})}function i(t){return Object(r["b"])({url:"/v1/platform/web/pm/productType/save",method:"post",data:t})}function u(t){return Object(r["b"])({url:"/v1/platform/web/pm/productType/update",method:"post",data:t})}function c(t){return Object(r["b"])({url:"/v1/platform/web/pm/productType/delete",method:"post",data:t})}function l(t){return Object(r["b"])({url:"/v1/platform/web/product/firmware/list",method:"post",data:t})}function d(t){return Object(r["b"])({url:"/v1/platform/web/product/firmware/setStatus",method:"post",data:t})}function s(t){return Object(r["b"])({url:"/v1/platform/web/product/firmwareVersion/add",method:"post",data:t})}function m(t){return Object(r["b"])({url:"/v1/platform/web/product/firmwareVersion/edit",method:"post",data:t})}function p(t){return Object(r["b"])({url:"/v1/platform/web/product/firmwareVersion/list",method:"post",data:t})}function f(t){return Object(r["b"])({url:"/v1/platform/web/product/firmwareVersion/onShelf",method:"post",data:t})}function v(t){return Object(r["b"])({url:"/v1/platform/web/product/firmwareVersion/unShelf",method:"post",data:t})}function y(t){return Object(r["b"])({url:"/v1/platform/web/product/firmware/add",method:"post",data:t})}function b(t){return Object(r["b"])({url:"/v1/platform/web/product/firmware/edit",method:"post",data:t})}function h(t){return Object(r["b"])({url:"/v1/platform/web/product/firmware/detail/"+t,method:"get"})}function q(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/get",method:"post",data:t})}function w(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/get",method:"get",params:t})}function k(t){return Object(r["b"])({url:"/v1/platform/web/pm/thingModel/getStandard",method:"get",params:t})}function g(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/status",method:"post",data:t})}function O(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/delete",method:"post",data:t})}function _(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/save",method:"post",data:t})}function P(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/update",method:"post",data:t})}function x(t){return Object(r["b"])({url:"/v1/platform/web/product/module/list",method:"post",data:t})}function j(t){return Object(r["b"])({url:"/v1/platform/web/product/module/delete",method:"post",data:t})}function S(t){return Object(r["b"])({url:"/v1/platform/web/product/module/detail/".concat(t),method:"get"})}function C(t){return Object(r["b"])({url:"/v1/platform/web/product/module/add",method:"post",data:t})}function D(t){return Object(r["b"])({url:"/v1/platform/web/product/module/edit",method:"post",data:t})}function T(t){return Object(r["b"])({url:"/v1/platform/web/product/module/setStatus",method:"post",data:t})}function L(t){return Object(r["b"])({url:"/v1/platform/web/iot/activeDevice/platformList",method:"post",data:t})}function $(t){return Object(r["b"])({url:"/v1/platform/web/iot/activeDevice/count",method:"post",data:t})}function N(t){return Object(r["b"])({url:"/v1/platform/web/iot/activeDevice/logList",method:"post",data:t})}function I(t){return Object(r["b"])({url:"/v1/platform/web/iot/activeDevice/logCount",method:"post",data:t})}function V(t){return Object(r["b"])({url:"/v1/platform/web/open/product/funcList",method:"get",params:t})}function K(t){return Object(r["b"])({url:"/v1/platform/web/iot/activeDevice/detail/".concat(t),method:"get"})}function E(){return Object(r["b"])({url:"/v1/platform/web/pm/networkGuide/GetDefaultNetworkGuides",method:"get"})}function z(t){return Object(r["b"])({url:"/v1/platform/web/pm/controlPanel/list",method:"post",data:t})}function F(t){return Object(r["b"])({url:"/v1/platform/web/pm/controlPanel/delete/".concat(t),method:"post",data:{id:t}})}function R(t){return Object(r["b"])({url:"/v1/platform/web/pm/controlPanel/add",method:"post",data:t,timeout:3e5})}function Y(t){return Object(r["b"])({url:"/v1/platform/web/pm/controlPanel/edit",method:"post",data:t,timeout:3e5})}function A(t){return Object(r["b"])({url:"/v1/platform/web/pm/controlPanel/detail/".concat(t),method:"get"})}function M(t){return Object(r["b"])({url:"/v1/platform/web/pm/controlPanel/setStatus",method:"post",data:t})}function G(t){return Object(r["b"])({url:"/v1/platform/web/pm/product/resetProductThingModels",method:"get",params:t})}},e706:function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t._self._c;return e("a-card",{attrs:{bordered:!1}},[e("div",{staticClass:"table-page-search-wrapper"},[e("a-form",{attrs:{layout:"inline"}},[e("a-row",{attrs:{gutter:16}},[e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入设备ID",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.did,callback:function(e){t.$set(t.queryParam.query,"did",e)},expression:"queryParam.query.did"}})],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入开发者账号",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.developer,callback:function(e){t.$set(t.queryParam.query,"developer",e)},expression:"queryParam.query.developer"}})],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入设备名称",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.deviceName,callback:function(e){t.$set(t.queryParam.query,"deviceName",e)},expression:"queryParam.query.deviceName"}})],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入产品key",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.productKey,callback:function(e){t.$set(t.queryParam.query,"productKey",e)},expression:"queryParam.query.productKey"}})],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-select",{attrs:{placeholder:"是否激活",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.isActive,callback:function(e){t.$set(t.queryParam.query,"isActive",e)},expression:"queryParam.query.isActive"}},t._l(t.$DictList("active_status"),(function(a){return e("a-select-option",{key:a.key,attrs:{value:a.key}},[t._v(" "+t._s(a.label)+" ")])})),1)],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-select",{attrs:{placeholder:"是否在线",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.isOnline,callback:function(e){t.$set(t.queryParam.query,"isOnline",e)},expression:"queryParam.query.isOnline"}},t._l(t.$DictList("online_status"),(function(a){return e("a-select-option",{key:a.key,attrs:{value:a.key}},[t._v(" "+t._s(a.label)+" ")])})),1)],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("span",{staticClass:"table-page-search-submitButtons"},[e("a-button",{attrs:{icon:"search",type:"primary"},on:{click:t.query}},[t._v(" 查询 ")]),e("a-divider",{attrs:{type:"vertical"}}),e("a-button",{attrs:{icon:"redo"},on:{click:t.reset}},[t._v(" 重置 ")])],1)])],1)],1)],1),e("div",{staticClass:"table-operator"},[e("a-row",{attrs:{gutter:48,type:"flex",justify:"space-between",align:"middle"}},[e("a-col",{attrs:{md:18,sm:24}},[e("a-row",[e("a-col",{attrs:{md:4,sm:24}},[e("div",[t._v("设备总数："+t._s(t.deviceCount.deviceTotal))])]),e("a-col",{attrs:{md:4,sm:24}},[e("div",[t._v("已激活设备："+t._s(t.deviceCount.activeTotal))])]),e("a-col",{attrs:{md:4,sm:24}},[e("div",[t._v("当前在线设备："+t._s(t.deviceCount.onlineTotal))])])],1)],1),e("a-col",[e("a-button",{attrs:{icon:"cloud-download"},on:{click:t.exportDevice}},[t._v(" 导出 ")])],1)],1)],1),e("a-table",{attrs:{size:"small",rowKey:"rowKey","data-source":t.dataSource,columns:t.columns,loading:t.loading,pagination:t.pagination},on:{change:t.onChangePagination},scopedSlots:t._u([{key:"deviceName",fn:function(a,r){return[e("div",{staticClass:"link-a",on:{click:function(e){return t.linkToPage("/product/device/details/index","DeviceDetails",r)}}},[t._v(" "+t._s(r.deviceName)+" "),e("br"),t._v(t._s(r.did)+" ")])]}},{key:"productName",fn:function(a,r){return[e("div",[t._v(t._s(r.productName)+" "),e("br"),t._v(t._s(r.productKey))])]}},{key:"activeStatus",fn:function(e,a){return[t._v(" "+t._s(t.$DictName("active_status",a.activeStatus))+" ")]}},{key:"onlineStatus",fn:function(e,a){return[t._v(" "+t._s(t.$DictName("online_status",a.onlineStatus))+" ")]}},{key:"activatedTime",fn:function(a,r){return[r.activatedTime?e("div",[t._v(" "+t._s(t._f("momentFilter")(r.activatedTime))+" ")]):e("div",[t._v("无")])]}},{key:"account",fn:function(a,r){return[e("div",[t._v(t._s(r.account)+" "),e("br"),t._v(t._s(r.companyName))])]}},{key:"action",fn:function(a){return[e("a-button",{attrs:{type:"link",size:"small",icon:"unordered-list"},on:{click:function(e){return t.linkToPage("/product/device/log/index","DeviceLogDetails",a)}}},[t._v(" 日志 ")])]}}])}),e("export-device",{ref:"exportDevice"})],1)},n=[],o=(a("99af"),a("d3b7"),a("159b"),a("a467")),i=a("c4c8"),u={name:"DeviceList",components:{"export-device":o["default"]},data:function(){return{isInit:!1,pagination:{showSizeChanger:!0,current:1,pageSize:10,total:0,showTotal:function(t,e){return"总数:".concat(t," 当前:").concat(e[0],"-").concat(e[1])}},queryParam:{page:1,limit:10,query:{}},columns:[{title:"序号",width:"46px",customRender:function(t,e,a){return a+1}},{dataIndex:"deviceName",title:"设备名称",scopedSlots:{customRender:"deviceName"}},{dataIndex:"activeStatus",title:"是否激活",scopedSlots:{customRender:"activeStatus"}},{dataIndex:"onlineStatus",title:"是否在线",scopedSlots:{customRender:"onlineStatus"}},{dataIndex:"productName",title:"所属产品",scopedSlots:{customRender:"productName"}},{dataIndex:"deviceNature",title:"设备性质"},{dataIndex:"account",title:"开发者",width:"15%",scopedSlots:{customRender:"account"}},{dataIndex:"activatedTime",title:"首次激活时间",scopedSlots:{customRender:"activatedTime"},width:"104px"},{title:"操作",key:"action",align:"center",width:"70px",scopedSlots:{customRender:"action"}}],dataSource:[],loading:!1,deviceCount:{},account:""}},computed:{},watch:{$route:function(){this.account=this.$route.query.account||"",this.queryParam.query.developer=this.$route.query.account||"",this.queryParam=this.$deepClone(this.queryParam)}},created:function(){this.isInit=!0,this.account=this.$route.query.account||"",this.queryParam.query.developer=this.$route.query.account||"",this.queryList(),this.init()},activated:function(){this.isInit?this.isInit=!1:this.queryList()},methods:{init:function(){var t=this,e={};this.account&&(e={query:{developer:this.account}}),Object(i["x"])(e).then((function(e){0==e.code&&(t.deviceCount=e.data)}))},exportDevice:function(){this.$refs["exportDevice"].visible=!0},onChangePagination:function(t){this.queryParam.page=t.current,this.queryParam.limit=t.pageSize,this.pagination.current=t.current,this.pagination.pageSize=t.pageSize,this.queryList()},query:function(){this.queryParam.page=1,this.pagination.current=1,this.queryList()},reset:function(){this.queryParam={page:1,limit:this.queryParam.limit,query:{}},this.pagination.current=1,this.queryList()},queryList:function(){var t=this;this.loading=!0,Object(i["z"])(this.queryParam).then((function(e){0==e.code&&(e.data.list.forEach((function(t,e){t.rowKey=e})),t.dataSource=e.data.list,t.pagination.total=e.data.total)})).finally((function(){t.loading=!1}))},linkToPage:function(t,e,a){this.$routerPush({path:t,query:{deviceName:a.deviceName,did:a.did,productId:a.productId},name:e})}}},c=u,l=a("2877"),d=Object(l["a"])(c,r,n,!1,null,"16caa6be",null);e["default"]=d.exports}}]);