Vue.component('location-select', {
    template:
        '<div class="form-group location-select">'+
        '    <label>{{label}}:</label>'+
        '    <input ref="selector" type="text" class="form-control form-control-sm" v-model="initialCopy">'+
        '</div>',


    props: ['label'],

    data: function () {
      return {
        initialCopy: ""
      }
    },

    mounted: function () {
        var vm = this;

        var my_autoComplete = new autoComplete({
            selector: this.$refs.selector,
            minChars: 3,
            delay: 500,
            cache: false,
            source: function(term, suggest){
                axios.get('/app/search/location?term='+encodeURIComponent(term))
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
                return '<div class="autocomplete-suggestion" data-id="'+item.ID+'" data-val="'+item.Name+'" data-type="'+item.Type+'">'
                       + item.Name.replace(re, "<b>$1</b>") + 
                       '</div>';
            },
            onSelect: function(e, term, item){
                vm.initialCopy="";
                vm.$emit(
                    'type-selected', 
                    parseInt(item.getAttribute('data-id')),
                    item.getAttribute('data-val'),   
                    item.getAttribute('data-type'),   
                );
            }
        });

    },
})
