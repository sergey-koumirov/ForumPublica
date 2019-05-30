Vue.component('pagination', {
    template:
        '<nav>'+
        '<ul class="pagination pagination-sm" v-if="visible">'+
        '  <li class="page-item" v-for="n in pageCount" :class="{active: n==current}">'+
        '    <span class="page-link" v-if="n==current">{{n}}</span>'+
        '    <a :href="\'?page=\'+n" class="page-link" v-if="n!=current">{{n}}</a>'+
        '  </li>'+
        '</ul>'+
        '</nav>',

    props: ['total','current','pp'],

    data: function () {
      return {
          visible: false,
          pageCount: 0,           
      }
    },

    mounted: function () {
        this.onChange();
    },

    methods: {
        onChange: function () {
            var vm = this;
            vm.visible = vm.total/vm.pp > 1;
            vm.pageCount = Math.ceil(vm.total/vm.pp);
        }
    },

    watch: {
        total: function(){
            this.onChange();
        }
    }

})
