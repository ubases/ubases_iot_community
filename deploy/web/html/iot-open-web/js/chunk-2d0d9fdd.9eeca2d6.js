(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-2d0d9fdd"],{"6a62":function(t,a,e){"use strict";e.r(a);var l=function(){var t=this,a=t._self._c;return a("a-modal",{attrs:{title:t.$t("manufacture.detail.title"),width:800,visible:t.visible},on:{cancel:t.handleCancel}},[a("h3",[t._v(t._s(t.$t("manufacture.detail.baseMsg")))]),a("a-form-model",{staticClass:"regular-form",attrs:{model:t.data,"label-col":{span:7},"wrapper-col":{span:13}}},[a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.did")}},[t._v(" "+t._s(t.data.did)+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.username")}},[t._v(" "+t._s(t.data.userName)+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.password")}},[t._v(" "+t._s(t.data.passward)+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.sn")}},[t._v(" "+t._s(t.data.sn)+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.batchId")}},[t._v(" "+t._s(t.data.batchId)+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.productName")}},[t._v(" "+t._s(t.data.productName)+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.isActived")}},[t._v(" "+t._s(t.$DictName("active_status",t.data.activeStatus))+" ")]),a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.createTime")}},[t._v(" "+t._s(t._f("momentFilter")(t.data.createdAt))+" ")])],1),a("h3",[t._v(t._s(t.$t("manufacture.detail.exportDetail")))]),a("a-form-model",{staticClass:"regular-form",attrs:{model:t.data,"label-col":{span:7},"wrapper-col":{span:13}}},[a("a-form-model-item",{attrs:{label:t.$t("manufacture.columns.exportTimes")}},[t._v(" "+t._s(t.data.exportCount)+" ")]),t._l(t.data.exportList||[],(function(e,l){return a("a-form-model-item",{key:l,attrs:{label:t.$t("manufacture.detail.exportTime")}},[t._v(" "+t._s(e.time)+" ")])}))],2),a("section",{staticClass:"ant-modal-footer",attrs:{slot:"footer"},slot:"footer"},[a("a-button",{on:{click:t.handleCancel}},[t._v(t._s(t.$t("public.close.text")))])],1)],1)},s=[],o={name:"ManufactureDetail",props:{visible:{type:Boolean,default:!1},data:{type:Object,default:function(){}}},data:function(){return{}},created:function(){},methods:{handleCancel:function(){this.$emit("handleClose")}}},r=o,m=e("2877"),n=Object(m["a"])(r,l,s,!1,null,"8d1bcea4",null);a["default"]=n.exports}}]);