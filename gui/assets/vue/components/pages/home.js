var HomeC = {
  template: vuedata.data.pages.home,
  props:{
    vbcd:Object,
    vicons:Object,
    vlng:Object,
  },
  data () {
    return {
      address:"",
      label:"",
      amount:0,
      transactions: blockchaindata.data.listtransactions,
      isPaginated: true,
      isPaginationSimple: false,
      defaultSortDirection: 'asc',
      currentPage: 1,
      perPage: 5,
    }
  },
  methods: {

  },
}

