(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2492efc7"],{"61f7":function(e,t,r){"use strict";r.d(t,"e",(function(){return n})),r.d(t,"k",(function(){return a})),r.d(t,"g",(function(){return u})),r.d(t,"q",(function(){return s})),r.d(t,"f",(function(){return p})),r.d(t,"r",(function(){return m})),r.d(t,"i",(function(){return f})),r.d(t,"j",(function(){return g})),r.d(t,"d",(function(){return b})),r.d(t,"b",(function(){return v})),r.d(t,"c",(function(){return h})),r.d(t,"m",(function(){return w})),r.d(t,"h",(function(){return j})),r.d(t,"o",(function(){return C})),r.d(t,"n",(function(){return x})),r.d(t,"p",(function(){return y})),r.d(t,"l",(function(){return k})),r.d(t,"a",(function(){return O}));r("99af");var n=/^[A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4}$/,a=/^[1][3,4,5,6,7,8,9][0-9]{9}$/,o=/^[A-Za-z0-9]+$/,l=/^([0-9]\d|[0-9])(.([0-9]\d|\d)){2}$/,i=/^([0-9]\d|[0-9])(.([0-9]\d|\d))(.([0-9]\d|\d)[A-Za-z0-9_]{0,})$/,u=/^[A-Za-z0-9_]+$/,d=/^[\a-\z\A-\Z0-9\.\,\?\<\>\。\，\-\——\=\;\@\！\!\+]+$/g,c=/^[A-Za-z0-9\u4e00-\u9fa5]+$/,s=/^[A-Za-z0-9\s\u4e00-\u9fa5]+$/,p=/^[A-Za-z0-9\s]+$/,m=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:c,message:"限制中文,英文和数字",trigger:t}]},f=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:/^[\a-\z\A-\Z0-9\ \.\,\?\<\>\。\，\-\——\=\;\@\！\!\+]+$/g,message:"限制英文,数字和部分常用符号",trigger:t}]},g=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:o,message:"只能输入英文和数字不能使用空格,符号",trigger:t}]},b=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"",r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"";return[{type:t,required:!0,message:e,trigger:r}]},v=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:0,t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:50,r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:"字符长度0-50",n=arguments.length>3&&void 0!==arguments[3]&&arguments[3];return[{required:n,min:e,max:t,message:r}]},h=function(e){var t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:50,r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:1,n=arguments.length>3&&void 0!==arguments[3]?arguments[3]:"";return[{required:!0,message:e,trigger:n},{required:!0,min:r,max:t,message:"字符长度在".concat(r,"-").concat(t),trigger:n}]},w=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入手机号码",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:a,message:"请输入正确的手机号码(11位)",trigger:t}]},j=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入邮箱",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:n,message:"请输入正确格式规则的邮箱",trigger:t}]},C=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入版本号",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:l,message:"请输入格式xx.xx.xx的版本号",trigger:t}]},x=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入版本号",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,pattern:i,message:"请输入格式xx.xx.xx的版本号",trigger:t}]},y=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入wifi名",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:1,r=arguments.length>2&&void 0!==arguments[2]?arguments[2]:4,n=arguments.length>3&&void 0!==arguments[3]?arguments[3]:"";return[{required:!0,message:e,trigger:n},{required:!0,min:t,max:r,message:"字符长度在".concat(t,"-").concat(r),trigger:n},{required:!0,pattern:u,message:"请输入英文,数字,下划线的wifi名,不能使用空格和其他符号",trigger:n}]},k=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入密码",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:8,max:50,message:"字符长度在8-50",trigger:t},{required:!0,pattern:d,message:"请输入英文,数字或常用符号的密码",trigger:t},{required:!0,pattern:/^(?=.*?[a-z])(?=.*?[A-Z])(?=.*?\d).*$/g,message:"密码必须要有英文字母大小写和数字",trigger:t}]},O=function(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:"请输入账号",t=arguments.length>1&&void 0!==arguments[1]?arguments[1]:"";return[{required:!0,message:e,trigger:t},{required:!0,min:1,max:50,message:"字符长度在1-50",trigger:t},{required:!0,pattern:/^([A-Za-z\d]+([-_.][A-Za-z\d]+)*@([A-Za-z\d]+[-.])+[A-Za-z\d]{2,4})|([1][3,4,5,6,7,8,9][0-9]{9})$/,message:"请输入国内11位手机号或者邮箱",trigger:t}]}},a679:function(e,t,r){"use strict";r.r(t);var n=function(){var e=this,t=e._self._c;return t("a-modal",{attrs:{title:e.modalTitle,width:640,visible:e.visible,"confirm-loading":e.confirmLoading},on:{cancel:e.handleCancel},scopedSlots:e._u([{key:"footer",fn:function(){return[t("a-button",{key:"back",on:{click:e.handleCancel}},[e._v("取消")]),t("a-divider",{attrs:{type:"vertical"}}),t("a-button",{key:"submit",attrs:{type:"primary",loading:e.confirmLoading},on:{click:e.handleOk}},[e._v(" 确定 ")])]},proxy:!0}])},[t("a-spin",{attrs:{spinning:e.confirmLoading}},[t("a-form-model",{ref:"ruleForm",attrs:{model:e.form,rules:e.rules}},[t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"模板类型",prop:"tplType"}},[t("a-select",{attrs:{placeholder:"请选择模板类型"},model:{value:e.form.tplType,callback:function(t){e.$set(e.form,"tplType",t)},expression:"form.tplType"}},e._l(e.$DictList("tpl_notice_type"),(function(r){return t("a-select-option",{key:r.key,attrs:{value:r.key}},[e._v(" "+e._s(r.label)+" ")])})),1)],1),t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"通知方式",prop:"method"}},[t("a-select",{attrs:{placeholder:"请选择通知方式"},model:{value:e.form.method,callback:function(t){e.$set(e.form,"method",t)},expression:"form.method"}},e._l(e.$DictList("tpl_notice_method"),(function(r){return t("a-select-option",{key:r.key,attrs:{value:r.key}},[e._v(" "+e._s(r.label)+" ")])})),1)],1),2!=e.form.method?t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"短信服务供应商",prop:"smsSupplier"}},[t("a-select",{attrs:{placeholder:"请选择短信服务供应商"},model:{value:e.form.smsSupplier,callback:function(t){e.$set(e.form,"smsSupplier",t)},expression:"form.smsSupplier"}},e._l(e.$DictList("sms_supplier"),(function(r){return t("a-select-option",{key:r.key,attrs:{value:r.key}},[e._v(" "+e._s(r.label)+" ")])})),1)],1):e._e(),2==e.form.method?t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"邮件服务供应商",prop:"smsSupplier"}},[t("a-select",{attrs:{placeholder:"请选择邮件服务供应商"},model:{value:e.form.smsSupplier,callback:function(t){e.$set(e.form,"smsSupplier",t)},expression:"form.smsSupplier"}},e._l(e.$DictList("email_supplier"),(function(r){return t("a-select-option",{key:r.key,attrs:{value:r.key}},[e._v(" "+e._s(r.label)+" ")])})),1)],1):e._e(),t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"第三方模板编码",prop:"thirdparyCode"}},[t("a-input",{attrs:{placeholder:"请输入第三方模板编码"},model:{value:e.form.thirdparyCode,callback:function(t){e.$set(e.form,"thirdparyCode",t)},expression:"form.thirdparyCode"}})],1),t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"通知名称",prop:"tplName"}},[t("a-input",{attrs:{placeholder:"请输入通知名称"},model:{value:e.form.tplName,callback:function(t){e.$set(e.form,"tplName",t)},expression:"form.tplName"}})],1),t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"通知主题",prop:"tplSubject"}},[t("a-input",{attrs:{placeholder:"请输入通知主题"},model:{value:e.form.tplSubject,callback:function(t){e.$set(e.form,"tplSubject",t)},expression:"form.tplSubject"}})],1),t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"模板语种",prop:"lang"}},[t("a-select",{attrs:{placeholder:"请选择模板语种",options:e.$DictList("language_type")},model:{value:e.form.lang,callback:function(t){e.$set(e.form,"lang",t)},expression:"form.lang"}})],1),t("a-form-model-item",{attrs:{"label-col":e.labelCol,"wrapper-col":e.wrapperCol,label:"通知内容",prop:"tplContent"}},[t("input-textarea",{attrs:{placeholder:"请输入通知内容",maxLength:1e3,rows:5},on:{input:function(t){return e.$refs["ruleForm"].clearValidate(["tplContent"])}},model:{value:e.form.tplContent,callback:function(t){e.$set(e.form,"tplContent",t)},expression:"form.tplContent"}})],1)],1)],1)],1)},a=[],o=r("c7eb"),l=r("1da1"),i=(r("d3b7"),r("c621")),u=r("61f7"),d={data:function(){return{id:"",type:"",modalTitle:"",visible:!1,confirmLoading:!1,form:{},rules:{tplName:Object(u["c"])("请输入通知名称"),method:Object(u["d"])("请选择通知方式","number"),smsSupplier:Object(u["d"])("请选择短信服务商","any"),tplType:Object(u["d"])("请选择短信模板类型","any"),tplSubject:Object(u["c"])("请输入通知主题"),thirdparyCode:Object(u["c"])("请输入第三方模板编码"),tplContent:Object(u["d"])("请输入通知内容"),lang:Object(u["d"])("请选择模板语种")},labelCol:{xs:{span:24},sm:{span:5}},wrapperCol:{xs:{span:24},sm:{span:19}}}},created:function(){},methods:{init:function(e,t){var r=this;this.type=e,this.$refs["ruleForm"]&&this.$refs["ruleForm"].resetFields(),this.form={tplName:"",tplSubject:"",tplContent:""},"add"===this.type?(this.modalTitle="添加通知模板",this.form={method:1}):"edit"===this.type&&(this.modalTitle="编辑通知模板",this.id=t,this.confirmLoading=!0,Object(i["n"])(t).then((function(e){0===e.code&&(r.form=e.data)})).finally((function(){r.confirmLoading=!1}))),this.visible=!0},handleOk:function(){var e=this;this.$refs["ruleForm"].validate(function(){var t=Object(l["a"])(Object(o["a"])().mark((function t(r){var n;return Object(o["a"])().wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(r){t.next=2;break}return t.abrupt("return");case 2:if(e.confirmLoading=!0,!e.form.id){t.next=9;break}return t.next=6,Object(i["i"])(e.form);case 6:n=t.sent,t.next=13;break;case 9:return e.form.tplCode=e.form.thirdparyCode,t.next=12,Object(i["b"])(e.form);case 12:n=t.sent;case 13:e.toast(n),e.confirmLoading=!1,0===n.code&&(e.$emit("ok"),e.visible=!1);case 16:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}())},handleCancel:function(){this.visible=!1}}},c=d,s=r("2877"),p=Object(s["a"])(c,n,a,!1,null,"e2e31a74",null);t["default"]=p.exports},c621:function(e,t,r){"use strict";r.d(t,"o",(function(){return a})),r.d(t,"f",(function(){return o})),r.d(t,"b",(function(){return l})),r.d(t,"i",(function(){return i})),r.d(t,"n",(function(){return u})),r.d(t,"s",(function(){return d})),r.d(t,"u",(function(){return c})),r.d(t,"d",(function(){return s})),r.d(t,"k",(function(){return p})),r.d(t,"r",(function(){return m})),r.d(t,"m",(function(){return f})),r.d(t,"e",(function(){return g})),r.d(t,"t",(function(){return b})),r.d(t,"a",(function(){return v})),r.d(t,"h",(function(){return h})),r.d(t,"l",(function(){return w})),r.d(t,"q",(function(){return j})),r.d(t,"c",(function(){return C})),r.d(t,"j",(function(){return x})),r.d(t,"p",(function(){return y})),r.d(t,"g",(function(){return k}));var n=r("b775");function a(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/noticeTpl/list",method:"post",data:e})}function o(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/noticeTpl/delete",method:"post",data:e})}function l(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/noticeTpl/add",method:"post",data:e})}function i(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/noticeTpl/edit",method:"post",data:e})}function u(e){return Object(n["b"])({url:"/v1/platform/web/template/noticeTpl/detail/".concat(e),method:"get"})}function d(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/testCaseTpl/list",method:"post",data:e})}function c(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/testCaseTpl/setStatus",method:"post",data:e})}function s(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/testCaseTpl/add",method:"post",data:e,timeout:3e5})}function p(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/testCaseTpl/edit",method:"post",data:e,timeout:3e5})}function m(e){return Object(n["b"])({url:"/v1/platform/web/template/testCaseTpl/detail/".concat(e),method:"get"})}function f(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/documentTpl/list",method:"post",data:e})}function g(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/documentTpl/delete",method:"post",data:e})}function b(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/documentTpl/setStatus",method:"post",data:e})}function v(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/documentTpl/add",method:"post",data:e})}function h(){var e=arguments.length>0&&void 0!==arguments[0]?arguments[0]:{};return Object(n["b"])({url:"/v1/platform/web/template/documentTpl/edit",method:"post",data:e})}function w(e){return Object(n["b"])({url:"/v1/platform/web/template/documentTpl/detail/".concat(e),method:"get"})}function j(e){return Object(n["b"])({url:"/v1/platform/web/template/messageTpl/list",method:"post",data:e})}function C(e){return Object(n["b"])({url:"/v1/platform/web/template/messageTpl/add",method:"post",data:e})}function x(e){return Object(n["b"])({url:"/v1/platform/web/template/messageTpl/edit",method:"post",data:e})}function y(e){return Object(n["b"])({url:"/v1/platform/web/template/messageTpl/detail/".concat(e),method:"get"})}function k(e){return Object(n["b"])({url:"/v1/platform/web/template/messageTpl/delete",method:"post",data:e})}}}]);