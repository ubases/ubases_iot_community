(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-09737c3e"],{"7e1e":function(t,e,a){"use strict";a.d(e,"h",(function(){return n})),a.d(e,"i",(function(){return i})),a.d(e,"d",(function(){return o})),a.d(e,"c",(function(){return u})),a.d(e,"e",(function(){return l})),a.d(e,"g",(function(){return c})),a.d(e,"f",(function(){return s})),a.d(e,"b",(function(){return d})),a.d(e,"a",(function(){return p}));var r=a("b775");function n(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/overview/accumulate",method:"get",params:t})}function i(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/overview/today",method:"get",params:t})}function o(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/developer/list",method:"post",data:t})}function u(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/developer/detail",method:"get",params:t})}function l(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/developer/total",method:"get",params:t})}function c(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/device/total",method:"get",params:t})}function s(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/deviceFault/list",method:"post",data:t})}function d(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/app/list",method:"post",data:t})}function p(){var t=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(r["b"])({url:"/v1/platform/web/data/pm/app/detail",method:"get",params:t})}},"9c99":function(t,e,a){"use strict";a.r(e);var r=function(){var t=this,e=t._self._c;return e("a-card",{attrs:{bordered:!1}},[e("div",{staticClass:"table-page-search-wrapper"},[e("a-form",{attrs:{layout:"inline"}},[e("a-row",{attrs:{gutter:48}},[e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入开发者账号",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.userName,callback:function(e){t.$set(t.queryParam.query,"userName",e)},expression:"queryParam.query.userName"}})],1)],1),e("a-col",t._b({},"a-col",t.aColFlex,!1),[e("a-form-item",[e("a-input",{attrs:{placeholder:"请输入APP名称",allowClear:!0},nativeOn:{keyup:function(e){return!e.type.indexOf("key")&&t._k(e.keyCode,"enter",13,e.key,"Enter")?null:t.query.apply(null,arguments)}},model:{value:t.queryParam.query.appName,callback:function(e){t.$set(t.queryParam.query,"appName",e)},expression:"queryParam.query.appName"}})],1)],1),e("a-col",{attrs:{md:4,sm:24}},[e("span",{staticClass:"table-page-search-submitButtons"},[e("a-button",{attrs:{icon:"search",type:"primary"},on:{click:t.query}},[t._v(" 查询 ")]),e("a-divider",{attrs:{type:"vertical"}}),e("a-button",{attrs:{icon:"redo"},on:{click:t.reset}},[t._v(" 重置 ")])],1)])],1)],1)],1),e("a-table",{attrs:{size:"small",rowKey:"appId","data-source":t.dataSource,columns:t.columns,loading:t.loading,pagination:t.pagination},on:{change:t.onChangePagination},scopedSlots:t._u([{key:"appName",fn:function(a,r){return[e("div",{staticClass:"link-a",on:{click:function(e){return t.linkToPage("/data-center/app-data/details",r)}}},[t._v(" "+t._s(r.appName)+" ")])]}}])})],1)},n=[],i=(a("99af"),a("d3b7"),a("7e1e")),o={name:"AppData",data:function(){return{isInit:!1,pagination:{showSizeChanger:!0,current:1,pageSize:10,total:0,showTotal:function(t,e){return"总数:".concat(t," 当前:").concat(e[0],"-").concat(e[1])}},queryParam:{page:1,limit:10,query:{}},columns:[{title:"序号",width:"80px",customRender:function(t,e,a){return a+1}},{dataIndex:"appName",title:"APP名称",scopedSlots:{customRender:"appName"}},{dataIndex:"developerId",title:"开发者账号"},{dataIndex:"registerUserTotal",title:"注册用户"},{dataIndex:"acitveUserTotal",title:"近七天活跃用户"},{dataIndex:"version",title:"最新版本"},{dataIndex:"verTotal",title:"故障累计"},{dataIndex:"feedbackQuantity",title:"用户反馈"}],dataSource:[],loading:!1}},computed:{},created:function(){this.isInit=!0,this.queryList()},activated:function(){this.isInit?this.isInit=!1:this.queryList()},methods:{exportDevice:function(){this.$refs["exportDevice"].visible=!0},onChangePagination:function(t){this.queryParam.page=t.current,this.queryParam.limit=t.pageSize,this.pagination.current=t.current,this.pagination.pageSize=t.pageSize,this.queryList()},query:function(){this.queryParam.page=1,this.pagination.current=1,this.queryList()},reset:function(){this.queryParam={page:1,limit:this.queryParam.limit,query:{}},this.pagination.current=1,this.queryList()},queryList:function(){var t=this;this.loading=!0,Object(i["b"])(this.queryParam).then((function(e){0===e.code&&(t.dataSource=e.data.list,t.pagination.total=e.data.total)})).finally((function(){t.loading=!1}))},linkToPage:function(t,e){this.$routerPush({path:t,name:"AppDataDetails",query:{id:e.appId}})}}},u=o,l=a("2877"),c=Object(l["a"])(u,r,n,!1,null,"dcd1525a",null);e["default"]=c.exports}}]);