(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-3068fc9b"],{"2a8c":function(e,t,r){"use strict";r.r(t);var n=function(){var e=this,t=e._self._c;return t("div",[t("a-card",{attrs:{bordered:!1}},[t("div",{staticClass:"table-operator"},[t("a-button",{attrs:{type:"primary",icon:"plus"},on:{click:function(t){return e.handleSave("add")}}},[e._v(" "+e._s(e.$t("public.add"))+" ")])],1),t("a-table",{attrs:{size:"small",rowKey:"id","data-source":e.configList,columns:e.columns,loading:e.loading,pagination:!1},scopedSlots:e._u([{key:"action",fn:function(r,n){return[t("a-button",{attrs:{type:"link",size:"small",icon:"edit"},on:{click:function(t){return e.handleSave("edit",n)}}},[e._v(" "+e._s(e.$t("public.edit"))+" ")]),t("a-divider",{attrs:{type:"vertical"}}),t("a-button",{attrs:{type:"link",size:"small",icon:"delete"},on:{click:function(t){return e.handleDelete(n)}}},[e._v(" 删除 ")])]}}])}),t("a-modal",{attrs:{title:e.title,width:640,visible:e.visible,"confirm-loading":e.confirmLoading},on:{ok:e.handleOk,cancel:e.handleCancel},scopedSlots:e._u([{key:"footer",fn:function(){return[t("a-button",{key:"back",on:{click:e.handleCancel}},[e._v("取消")]),t("a-divider",{attrs:{type:"vertical"}}),t("a-button",{key:"submit",attrs:{type:"primary",loading:e.confirmLoading},on:{click:e.handleOk}},[e._v(" 确定 ")])]},proxy:!0}])},[t("a-spin",{attrs:{spinning:e.confirmLoading}},[t("a-form-model",{ref:"ruleForm",attrs:{model:e.form,rules:e.rules}},[t("a-row",[t("a-col",{attrs:{span:24}},[t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"参数描述",prop:"desc"}},[t("a-input",{attrs:{placeholder:"请输入参数描述"},model:{value:e.form.desc,callback:function(t){e.$set(e.form,"desc",t)},expression:"form.desc"}})],1)],1)],1),t("a-row",[t("a-col",{attrs:{span:24}},[t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"参数编码",prop:"code"}},[t("a-input",{attrs:{placeholder:"请输入参数编码"},model:{value:e.form.code,callback:function(t){e.$set(e.form,"code",t)},expression:"form.code"}})],1)],1)],1),t("a-row",[t("a-col",{attrs:{span:24}},[t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"参数值",prop:"value"}},[t("a-input",{attrs:{placeholder:"请输入参数值"},model:{value:e.form.value,callback:function(t){e.$set(e.form,"value",t)},expression:"form.value"}})],1)],1)],1)],1)],1)],1)],1)],1)},i=[],a=r("c7eb"),o=r("1da1"),d=(r("d3b7"),r("b775"));function s(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(d["b"])({url:"/v1/platform/web/config/systemConfig/list",method:"post",data:e})}function l(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(d["b"])({url:"/v1/platform/web/config/systemConfig/add",method:"post",data:e})}function u(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(d["b"])({url:"/v1/platform/web/config/systemConfig/edit",method:"post",data:e})}function c(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(d["b"])({url:"/v1/platform/web/config/systemConfig/delete",method:"post",data:e})}var g=r("61f7"),f={name:"Configuration",components:{},data:function(){return{configList:[],columns:[{title:"序号",width:"80px",customRender:function(e,t,r){return r+1}},{dataIndex:"desc",title:"参数描述",ellipsis:!0},{dataIndex:"code",title:"参数编码",ellipsis:!0},{dataIndex:"value",title:"参数值",ellipsis:!0},{title:this.$t("public.action"),key:"action",align:"center",width:"180px",scopedSlots:{customRender:"action"}}],loading:!1,visible:!1,title:"",confirmLoading:!1,form:{},rules:{desc:Object(g["c"])("请输入参数描述",100,1),code:Object(g["c"])("请输入参数编码",50,1),value:Object(g["c"])("请输入参数值",500,1)},labelCol:{xs:{span:24},sm:{span:4}},wrapperCol:{xs:{span:24},sm:{span:20}}}},computed:{},mounted:function(){this.init()},methods:{init:function(){var e=this;this.loading=!0,s().then((function(t){0===t.code&&(e.configList=t.data.list)})).finally((function(){e.loading=!1}))},handleDelete:function(e){var t=this;this.$confirm({title:"确定是否删除",content:"删除后，将不可恢复",onOk:function(){c({id:e.id}).then((function(e){0===e.code&&t.init(),t.toast(e)}))}})},handleSave:function(e,t){"add"===e?(this.title="添加平台配置项",this.form={}):"edit"===e&&(this.title="编辑平台配置项",this.form=t),this.visible=!0},handleOk:function(){var e=this;this.$refs["ruleForm"].validate(function(){var t=Object(o["a"])(Object(a["a"])().mark((function t(r){var n;return Object(a["a"])().wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(r){t.next=2;break}return t.abrupt("return");case 2:if(e.confirmLoading=!0,e.form.id){t.next=9;break}return t.next=6,l(e.form);case 6:n=t.sent,t.next=12;break;case 9:return t.next=11,u(e.form);case 11:n=t.sent;case 12:e.toast(n),e.confirmLoading=!1,0===n.code&&(e.visible=!1,e.init());case 15:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}())},handleCancel:function(){this.visible=!1}}},m=f,p=r("2877"),v=Object(p["a"])(m,n,i,!1,null,"19cb7cda",null);t["default"]=v.exports},"61f7":function(e,t,r){"use strict";r.d(t,"e",(function(){return n})),r.d(t,"k",(function(){return i})),r.d(t,"g",(function(){return s})),r.d(t,"q",(function(){return c})),r.d(t,"f",(function(){return g})),r.d(t,"r",(function(){return f})),r.d(t,"i",(function(){return m})),r.d(t,"j",(function(){return p})),r.d(t,"d",(function(){return v})),r.d(t,"b",(function(){return h})),r.d(t,"c",(function(){return b})),r.d(t,"m",(function(){return x})),r.d(t,"h",(function(){return q})),r.d(t,"o",(function(){return k})),r.d(t,"n",(function(){return w})),r.d(t,"p",(function(){return y})),r.d(t,"l",(function(){return $})),r.d(t,"a",(function(){return z}));r("99af");var n=/^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$/,i=/^[1][3,4,5,6,7,8,9][0-9]{9}$/,a=/^[A-Za-z0-9]+$/,o=/^([0-9]\d|[0-9])(.([0-9]\d|\d)){2}$/,d=/^([0-9]\d|[0-9])(.([0-9]\d|\d))(.([0-9]\d|\d)[A-Za-z0-9_]{0,})$/,s=/^[A-Za-z0-9_]+$/,l=/^[\a-\z\A-\Z0-9\.\,\?\<\>\。\，\-\——\=\;\@\！\!\+]+$/g,u=/^[A-Za-z0-9\u4e00-\u9fa5]+$/,c=/^[A-Za-z0-9\s\u4e00-\u9fa5]+$/,g=/^[A-Za-z0-9\s]+$/,f=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:u,message:"限制中文,英文和数字",trigger:t}]},m=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:/^[\a-\z\A-\Z0-9\ \.\,\?\<\>\。\，\-\——\=\;\@\！\!\+]+$/g,message:"限制英文,数字和部分常用符号",trigger:t}]},p=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:a,message:"只能输入英文和数字不能使用空格,符号",trigger:t}]},v=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"",r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";return[{type:t,required:!0,message:e,trigger:r}]},h=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:0,t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:50,r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"字符长度0-50",n=arguments.length>3&&void 0!==arguments[3]&&arguments[3];return[{required:n,min:e,max:t,message:r}]},b=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:50,r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:1,n=arguments.length>3&&void 0!==arguments[3]?arguments[3]:"";return[{required:!0,message:e,trigger:n},{required:!0,min:r,max:t,message:"字符长度在".concat(r,"-").concat(t),trigger:n}]},x=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入手机号码",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:i,message:"请输入正确的手机号码(11位)",trigger:t}]},q=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入邮箱",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:n,message:"请输入正确格式规则的邮箱",trigger:t}]},k=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入版本号",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:o,message:"请输入格式xx.xx.xx的版本号",trigger:t}]},w=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入版本号",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:d,message:"请输入格式xx.xx.xx的版本号",trigger:t}]},y=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入wifi名",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:1,r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:4,n=arguments.length>3&&void 0!==arguments[3]?arguments[3]:"";return[{required:!0,message:e,trigger:n},{required:!0,min:t,max:r,message:"字符长度在".concat(t,"-").concat(r),trigger:n},{required:!0,pattern:s,message:"请输入英文,数字,下划线的wifi名,不能使用空格和其他符号",trigger:n}]},$=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入密码",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:8,max:50,message:"字符长度在8-50",trigger:t},{required:!0,pattern:l,message:"请输入英文,数字或常用符号的密码",trigger:t},{required:!0,pattern:/^(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d).*$/g,message:"密码必须要有英文字母大小写和数字",trigger:t}]},z=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入账号",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:/^([A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4})|([1][3,4,5,6,7,8,9][0-9]{9})$/,message:"请输入国内11位手机号或者邮箱",trigger:t}]}}}]);