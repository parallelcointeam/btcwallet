var SendC = {
  template: vuedata.data.pages.send,
  props:{
    vicons:Object,
    vlng:Object,

  },
  data () {
    return {
      address:"",
      label:"",
      amount:0,
      transactions: blockchaindata.data.listallsendtransactions,
      isPaginated: true,
      isPaginationSimple: false,
      defaultSortDirection: 'asc',
      currentPage: 1,
      perPage: 5,
    }
  },
  components: {
},
  methods: {
    sendForm: function() {
      sendtoaddress.sendDUO(this.address, this.label, this.amount);
     },
    },

}

