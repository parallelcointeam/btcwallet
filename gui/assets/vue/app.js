// Vue.use(Buefy);
// Vue.use(EasyBar);
// Vue.use(VueLayers);

var app = new Vue({
  el: '#app',
  data () {
    return {
    vdt: vuedata,
    bcd: blockchaindata,
    lng: language,
    // rpc: rpchandlers,
    // vpage: vdt.data.pages.home,
    timer: '',
    // component: Home,
    // component: HomeC,
    updateAvailable: false,
  }
},
components: {
  // HomeC,
  // SendC,
},
created: function() {
  // this.ref();
  // this.lang();
  // this.timer = setInterval(this.ref, 500)
  // this.timer = setInterval(this.lang, 500)
},
methods: {
  // processForm: function() {
  //   console.log({ name: this.name, email: this.email });
  //   alert('Processing');
  // },
  // rpc: function() { rpchandlers},
  swapComponent: function(component){this.component = component;},
  ref: function() { blockchaindata.getInfoData(); },
  lang: function() { language.languageData(vuedata.data.config.lang); },
  cancelAutoUpdate: function() { clearInterval(this.timer) }
},
beforeDestroy() {
clearInterval(this.timer)
}
});
