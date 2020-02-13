var transactions = new Vue({
    el: '#transactions',

    data: {
      data: {},
      isLoading: false,
    },

    created: function () {
      new Clipboard('.copy-item');
    },

    beforeMount: function () {
        this.data = JSON.parse(this.$el.attributes['transactions'].value);
    },

    methods: {
        toggleSummary: (record)=>{
            record.InSummary = record.InSummary ?  false : true;
        },
    },

    computed: {
        summary: function(){
            var result = 0;
            this.data.Records.forEach((el)=>{
                if(el.InSummary){
                    result = result + el.Quantity * el.Price;
                }
            });
            return result;
        },
    },
});


