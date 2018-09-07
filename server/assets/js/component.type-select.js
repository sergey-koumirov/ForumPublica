Vue.component('type-select', {
    template:
        '<div class="form-group type-select">'+
        '    <label v-if="mode!=\'compact\'">{{label}}:</label>'+
        '    <input ref="selector" type="text" class="form-control form-control-sm" :value="initialCopy">'+
        '</div>',


    props: ['label','initial','mode','filter'],

    data: function () {
      return {
        initialCopy: ""
      }
    },

    mounted: function () {
        var vm = this,
            filter = vm.filter||'item_type'
            vm.initialCopy = vm.initial;

        var my_autoComplete = new autoComplete({
            selector: this.$refs.selector,
            minChars: 3,
            delay: 500,
            cache: false,
            source: function(term, suggest){
                axios.get('/app/search/'+filter+'?term='+encodeURIComponent(term))
                    .then(function (r) {
                        if(!!r.data){
                            suggest(r.data);
                        }
                    })
                    .catch(function (error) {
                        console.log(error);
                    });
            },
            renderItem: function (item, search){
                search = search.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
                var re = new RegExp("(" + search.split(' ').join('|') + ")", "gi");
                return '<div class="autocomplete-suggestion" data-id="'+item.Id+'" data-val="'+item.Name+'">'+item.Name.replace(re, "<b>$1</b>")+'</div>';
            },
            onSelect: function(e, term, item){
                vm.$emit('type-selected', parseInt(item.getAttribute('data-id'))  );
                vm.initialCopy="";
            }
        });

    },
})
