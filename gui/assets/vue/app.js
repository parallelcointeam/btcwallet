Vue.use(Buefy);
Vue.use(VueVirtualScroller);
Vue.use(VueLayers);
Vue.use(VueTerminal);
 
var app = new Vue({
  el: '#app',
  data () {
    return {
    vdt: vuedata,
    bcd: blockchaindata,
    lng: language,
    ab: addressbook,
    rpc:rpcinterface,
    // rpc: rpchandlers,
    // vpage: vdt.data.pages.home,
    timer: '',
    // component: Home,
    component: HomeC,
    updateAvailable: false,
  }
},
components: {
  HomeC,
  // SendC,
},
created: function() {
  this.ref();
  this.adrbk();
  this.lang();
  this.getBlockCount();
  this.getInfo();
  this.timer = setInterval(this.ref, 1500)
  this.timer = setInterval(this.lang, 1500)
  this.timer = setInterval(this.getBlockCount, 500)
  this.timer = setInterval(this.getInfo, 500)
  // this.timer = setInterval(this.adrbk, 500)
},

methods: {
  // processForm: function() {
  //   console.log({ name: this.name, email: this.email });
  //   alert('Processing');
  // },
  // rpc: function() { rpchandlers},
  danger() {
    this.$snackbar.open({
        duration: 8000,
        message: rpcinterface.data.MSG,
        type: 'is-danger',
        position: 'is-bottom',
        actionText: 'close',
        queue: false,
        onAction: () => {
            this.$toast.open({
                message: 'Closed',
                queue: false
            })
        }
    })
},
  swapComponent: function(component){this.component = component;},
  ref: function() { blockchaindata.getInfoData(); },
  adrbk: function() { addressbook.addressBookData(); },
  getBlockCount: function() { rpcinterface.getBlockCount(); },
  getInfo: function() { rpcinterface.getInfo(); },
  lang: function() { language.languageData(vuedata.data.config.lang); },
  cancelAutoUpdate: function() { clearInterval(this.timer) }
},
beforeDestroy() {
clearInterval(this.timer)
}
});

  