import{k as f,l as m,t as _,i as t,e as a,m as C,n as H,b as I,s as v,F}from"./index.js";import{H as M,M as N}from"./HistShow.js";var P=_('<div class="card border-primary"><div class="card-header d-flex justify-content-between"></div><div class=card-body><table class="table table-striped table-hover"><tbody></tbody>'),j=_("<tr><td class=opacity-50 style=width:2em;>.</td><td><a></a><br><a></a></td><td>");function D(){const $=localStorage.getItem("histShow");return f(+$),(m()===0||isNaN(m()))&&f(200),(()=>{var d=P(),l=d.firstChild,u=l.nextSibling,S=u.firstChild,y=S.firstChild;return t(l,a(C,{}),null),t(l,a(M,{name:"histShow"}),null),t(y,a(F,{each:H,children:(e,g)=>(()=>{var n=j(),i=n.firstChild,w=i.firstChild,o=i.nextSibling,s=o.firstChild,p=s.nextSibling,c=p.nextSibling,x=o.nextSibling;return t(i,()=>g()+1,w),t(s,()=>e.Name),t(c,()=>e.IP),t(x,a(N,{get mac(){return e.Mac}})),I(r=>{var h="/host/"+e.ID,b="http://"+e.IP;return h!==r.e&&v(s,"href",r.e=h),b!==r.t&&v(c,"href",r.t=b),r},{e:void 0,t:void 0}),n})()})),d})()}export{D as default};
