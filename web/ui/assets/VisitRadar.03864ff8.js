import{A as r,bZ as s,z as i,S as d,a0 as n,B as u,a1 as l,a5 as m,H as p,X as f}from"./vendor.52d192f7.js";/* empty css                *//* empty css               *//* empty css               */import{u as c}from"./useECharts.5dff2d35.js";import{_ as h}from"./index.c0a3bc85.js";const x=r({components:{Card:s},props:{loading:Boolean,width:{type:String,default:"100%"},height:{type:String,default:"300px"}},setup(e){const t=i(null),{setOptions:a}=c(t);return d(()=>e.loading,()=>{e.loading||a({legend:{bottom:0,data:["\u8BBF\u95EE","\u8D2D\u4E70"]},tooltip:{},radar:{radius:"60%",splitNumber:8,indicator:[{text:"\u7535\u8111",max:100},{text:"\u5145\u7535\u5668",max:100},{text:"\u8033\u673A",max:100},{text:"\u624B\u673A",max:100},{text:"Ipad",max:100},{text:"\u8033\u673A",max:100}]},series:[{type:"radar",symbolSize:0,areaStyle:{shadowBlur:0,shadowColor:"rgba(0,0,0,.2)",shadowOffsetX:0,shadowOffsetY:10,opacity:1},data:[{value:[90,50,86,40,50,20],name:"\u8BBF\u95EE",itemStyle:{color:"#b6a2de"}},{value:[70,75,70,76,20,85],name:"\u8D2D\u4E70",itemStyle:{color:"#5ab1ef"}}]}]})},{immediate:!0}),{chartRef:t}}});function g(e,t,a,B,y,C){const o=n("Card");return u(),l(o,{title:"\u8F6C\u5316\u7387",loading:e.loading},{default:m(()=>[p("div",{ref:"chartRef",style:f({width:e.width,height:e.height})},null,4)]),_:1},8,["loading"])}var v=h(x,[["render",g]]);export{v as default};
