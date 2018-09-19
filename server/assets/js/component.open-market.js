Vue.component('open-market', {
    template:
      '<div class="popup">'+
      '  <div class="close" v-on:click="Close()">x&nbsp;</div>'+
      '  <div class="content">'+
      '    <ul>'+
      '      <li class="pointer" v-for="c in chars" v-on:click="OpenMarket(c.Id)">'+
      '        <span class="oi oi-external-link"></span>{{c.Name}}'+
      '      </li>'+
      '    </ul>'+
      '  </div>'+
      '</div>',

    props: ['chars'],

    data: function () {
      return {
        typeId: 0
      }
    },

    mounted: function () {
      var vm = this;
      vm.$root.$on('open-market', function(typeId,x,y){
          $(vm.$el).show().offset({left: x, top: y});
          vm.typeId = typeId;
      });
    },

    methods: {
      OpenMarket: function(id){
        var vm = this;
        axios.post(
          "/app/ui/open_market",
          {TypeId: vm.typeId, CharacterId: id}
        ).then(function(response){
          vm.Close();
        }).catch(function(error){
          vm.Close();
          console.log(error);
        });
      },

      Close: function(){
        $(this.$el).hide();
      },
    },

})
