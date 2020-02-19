Vue.component('location-select', {
    template:
        '<div class="form-group location-select">'+
        '    <label>{{label}}:</label>'+
        '    <input v-show="selected.id==null" ref="selector" type="text" class="form-control form-control-sm" v-model="initialCopy" :disabled="characterId==\'\'">'+
        '    <img v-if="!!selected.type" class="icon-suggestion" width="16" height="16" :src="\'/assets/images/\'+selected.type+\'.png\'" />'+
        '    <span v-show="selected.id!=null">{{selected.text}}</span>'+
        '    <span v-show="selected.id!=null" class="oi oi-delete pointer" @click="resetSelected"></span>'+
        '</div>',


    props: ['label','filter','characterId'],

    data: function () {
      return {
        initialCopy: "",
        selected: {
            id: null,
            text: null,
            type: null
        }
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
                axios.get('/app/search/location?term='+encodeURIComponent(term)+'&filter='+(vm.filter||'')+'&cid='+vm.characterId)
                    .then(function (r) {
                        if(!!r.data){
                            suggest(r.data);
                        }
                    })
                    .catch(function (error) {
                        console.log(error.response);
                        Toastify({text: error.response.data.error, close: true, duration: -1}).showToast();
                    });
            },
            renderItem: function (item, search){
                search = search.replace(/[-\/\\^$*+?.()|[\]{}]/g, '\\$&');
                var re = new RegExp("(" + search.split(' ').join('|') + ")", "gi"),
                    result = '<div class="autocomplete-suggestion" data-id="'+item.ID+'" data-val="'+item.Name+'" data-type="'+item.Type+'">';
                if(item.Type == "solar_system" || item.Type == "station" || item.Type == "structure"){
                    result = result + '<img class="icon-suggestion" width="16" height="16" src="/assets/images/'+item.Type+'.png" />';
                }
                result = result + item.Name.replace(re, "<b>$1</b>") + '</div>';
                return result;
            },
            onSelect: function(e, term, item){
                vm.initialCopy="";
                vm.selected.id = parseInt(item.getAttribute('data-id'));
                vm.selected.text = item.getAttribute('data-val');
                vm.selected.type = item.getAttribute('data-type');

                vm.$emit(
                    'location-selected',
                    vm.selected.id,
                    vm.selected.text,   
                    vm.selected.type,   
                );
                
            }
        });

    },

    methods: {
        resetSelected: function(){
            this.selected = {
                id: null,
                text: null,
                type: null
            };
            
            this.$emit(
                'location-reset'
            );
        },
      },
});
