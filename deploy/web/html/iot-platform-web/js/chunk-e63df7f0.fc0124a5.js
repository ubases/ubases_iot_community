(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-e63df7f0"],{"7e1e":function(t,e,a){"use strict";a.d(e,"h",(function(){return r})),a.d(e,"i",(function(){return o})),a.d(e,"d",(function(){return i})),a.d(e,"c",(function(){return d})),a.d(e,"e",(function(){return c})),a.d(e,"g",(function(){return u})),a.d(e,"f",(function(){return l})),a.d(e,"b",(function(){return s})),a.d(e,"a",(function(){return p}));var n=a("b775");function r(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/overview/accumulate",method:"get",params:t})}function o(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/overview/today",method:"get",params:t})}function i(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/developer/list",method:"post",data:t})}function d(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/developer/detail",method:"get",params:t})}function c(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/developer/total",method:"get",params:t})}function u(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/device/total",method:"get",params:t})}function l(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/deviceFault/list",method:"post",data:t})}function s(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/app/list",method:"post",data:t})}function p(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/data/pm/app/detail",method:"get",params:t})}},b0dc:function(t,e,a){},cd00:function(t,e,a){"use strict";a.r(e);var n=function(){var t=this,e=t._self._c;return e("a-card",{attrs:{bordered:!1}},[e("a-page-header",{attrs:{title:"开发者数据"},on:{back:function(e){return t.$multiTab.closeCurrentPage()}}},[e("a-spin",{attrs:{spinning:t.confirmLoading}},[e("a-card",{attrs:{title:"账号信息",bordered:!1}},[e("a-descriptions",[e("a-descriptions-item",{attrs:{label:"平台主账号"}},[t._v(" "+t._s(t.developerInfo.account)+" ")]),e("a-descriptions-item",{attrs:{label:"企业"}},[t._v(" "+t._s(t.developerInfo.companyName)+" ")]),e("a-descriptions-item",{attrs:{label:"角色"}},[t._v(" "+t._s(t.developerInfo.roleName)+" ")])],1)],1),e("a-card",{attrs:{title:"已激活设备数量：累计".concat(t.developerInfo.activeDeviceTotal,"个"),bordered:!1}},[e("a-button",{attrs:{type:"primary"},on:{click:t.linkToPage}},[t._v(" 查看设备详细，请跳转到“设备数据” ")])],1),e("a-card",{attrs:{title:"已开发APP：".concat(t.developerInfo.appList.length,"个"),bordered:!1}},[e("a-table",{attrs:{size:"small",rowKey:"appId","data-source":t.developerInfo.appList,columns:t.columnsApp,pagination:!1},scopedSlots:t._u([{key:"devStatus",fn:function(e,a){return[t._v(" "+t._s(t.$DictName("oem_app_status",a.devStatus))+" ")]}}])})],1)],1)],1)],1)},r=[],o=(a("d3b7"),a("7e1e")),i={name:"DeveloperDataDetails",data:function(){return{two:{labelCol:{xs:{span:24},sm:{span:4}},wrapperCol:{xs:{span:24},sm:{span:20}}},id:"",developerInfo:{appList:[]},confirmLoading:!1,columnsApp:[{dataIndex:"appName",title:"APP名称"},{dataIndex:"devStatus",title:"开发状态",scopedSlots:{customRender:"devStatus"}},{dataIndex:"version",title:"最新版本号"},{dataIndex:"verTotal",title:"历史迭代版本"}]}},created:function(){this.init()},methods:{linkToPage:function(){this.$routerPush({path:"/product/device/index?account=".concat(this.developerInfo.account),name:"Device"})},init:function(){var t=this;this.id=this.$route.query.id,this.confirmLoading=!0,Object(o["c"])({userId:this.id}).then((function(e){0===e.code&&(t.developerInfo=e.data)})).finally((function(){t.confirmLoading=!1}))}}},d=i,c=(a("ef70"),a("2877")),u=Object(c["a"])(d,n,r,!1,null,"378c7a62",null);e["default"]=u.exports},ef70:function(t,e,a){"use strict";a("b0dc")}}]);