(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["account","chunk-2d0da574"],{"6ad7":function(e,t,i){"use strict";i.r(t);var n=function(){var e=this,t=e._self._c;return t("section",{staticClass:"page-wrap no-padding"},[t("a-page-header",{attrs:{title:e.$t("securitySet.title")}},[t("section",{staticClass:"content"},[t("a-form",{attrs:{form:e.form,"label-col":e.labelCol,"wrapper-col":e.wrapperCol}},[t("a-form-item",{attrs:{label:e.$t("securitySet.password.label")}},[e._v(" ******* "),t("a-button",{attrs:{type:"link"},on:{click:e.editPassword}},[e._v(e._s(e.$t("securitySet.editPassword.button")))])],1),t("a-form-item",{attrs:{label:e.$t("securitySet.avoidLogin.label")}},[e._v(" "+e._s(e.$DictName("avoid_login",e.securityTime)||e.$t("securitySet.avoidLogin.default"))+" "),t("a-button",{attrs:{type:"link"},on:{click:e.handleChange}},[e._v(e._s(e.$t("securitySet.change.button")))])],1)],1)],1)]),t("a-modal",{attrs:{title:e.$t("securitySet.changeLogin.title"),visible:e.visible,width:250,"confirm-loading":e.confirmLoading},on:{ok:e.handleOk,cancel:e.handleCancel}},[t("a-spin",{attrs:{spinning:e.confirmLoading}},[t("a-form-model",{ref:"loginForm",attrs:{model:e.loginForm,rules:e.rules,"wrapper-col":{span:24}}},[t("a-form-model-item",{attrs:{prop:"loginType"}},[t("a-select",{attrs:{options:e.$DictList("avoid_login")},model:{value:e.loginForm.loginType,callback:function(t){e.$set(e.loginForm,"loginType",t)},expression:"loginForm.loginType"}})],1)],1)],1)],1)],1)},a=[],o=i("c7eb"),r=i("1da1"),s=i("5530"),l=i("2f62"),c={name:"SecuritySet",data:function(){return{labelCol:{xs:{span:24},sm:{span:3}},wrapperCol:{xs:{span:24},sm:{span:10}},form:{},visible:!1,confirmLoading:!1,loginForm:{loginType:""},rules:{loginType:[{required:!0,message:this.$t("securitySet.loginType.rules"),trigger:"blur"}]}}},created:function(){this.loginForm.loginType=this.securityTime},computed:Object(s["a"])({},Object(l["c"])(["securityTime"])),methods:Object(s["a"])(Object(s["a"])({},Object(l["d"])(["SET_SECURITY_TIME"])),{},{handleChange:function(){this.visible=!0},editPassword:function(){this.$router.push({name:"forgotPassword",query:{type:1}})},handleOk:function(){var e=this;this.$refs.loginForm.validate(function(){var t=Object(r["a"])(Object(o["a"])().mark((function t(i){return Object(o["a"])().wrap((function(t){while(1)switch(t.prev=t.next){case 0:if(i){t.next=2;break}return t.abrupt("return");case 2:return t.next=4,e.SET_SECURITY_TIME(e.loginForm.loginType);case 4:e.visible=!1;case 5:case"end":return t.stop()}}),t)})));return function(e){return t.apply(this,arguments)}}())},handleCancel:function(){this.visible=!1}})},u=c,p=i("2877"),d=Object(p["a"])(u,n,a,!1,null,"7ddc2648",null);t["default"]=d.exports}}]);