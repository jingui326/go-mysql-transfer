import{_ as k,G as b,a as h,W as v,e as R,Q as O,b as M,b0 as A,a4 as f,aU as B,aV as T,h as j}from"./index.c0a3bc85.js";import{A as $,c0 as x,be as D,a0 as l,B as F,D as N,k as o,a5 as i,ad as d,J as u,K as P,u as m}from"./vendor.52d192f7.js";const V=$({name:"SettingFooter",components:{CopyOutlined:x,RedoOutlined:D},setup(){const e=b(),{prefixCls:p}=h("setting-footer"),{t:s}=M(),{createSuccessModal:g,createMessage:r}=j(),C=v(),c=R(),t=O();function a(){const{isSuccessRef:n}=A(JSON.stringify(m(t.getProjectConfig),null,2));m(n)&&g({title:s("layout.setting.operatingTitle"),content:s("layout.setting.operatingContent")})}function S(){try{t.setProjectConfig(f);const{colorWeak:n,grayMode:y}=f;B(n),T(y),r.success(s("layout.setting.resetSuccess"))}catch(n){r.error(n)}}function _(){localStorage.clear(),t.resetAllState(),e.resetState(),C.resetState(),c.resetState(),location.reload()}return{prefixCls:p,t:s,handleCopy:a,handleResetSetting:S,handleClearAndRedo:_}}});function W(e,p,s,g,r,C){const c=l("CopyOutlined"),t=l("a-button"),a=l("RedoOutlined");return F(),N("div",{class:P(e.prefixCls)},[o(t,{type:"primary",block:"",onClick:e.handleCopy},{default:i(()=>[o(c,{class:"mr-2"}),d(" "+u(e.t("layout.setting.copyBtn")),1)]),_:1},8,["onClick"]),o(t,{color:"warning",block:"",onClick:e.handleResetSetting,class:"my-3"},{default:i(()=>[o(a,{class:"mr-2"}),d(" "+u(e.t("common.resetText")),1)]),_:1},8,["onClick"]),o(t,{color:"error",block:"",onClick:e.handleClearAndRedo},{default:i(()=>[o(a,{class:"mr-2"}),d(" "+u(e.t("layout.setting.clearBtn")),1)]),_:1},8,["onClick"])],2)}var I=k(V,[["render",W],["__scopeId","data-v-2d4de409"]]);export{I as default};
