Vue.component('pagination', {
    template:
        '<nav> {{total}} / {{current}} / {{pp}}'+
        '<ul class="pagination pagination-sm">'+
        '  <li class="page-item disabled">'+
        '    <span class="page-link">&lt;</span>'+
        '    <a href="?page=2" class="page-link">&gt;</a>'+
        '  </li>'+
        
        '  <li class="page-item active"><span class="page-link">1</span></li>'+
        '  <li class="page-item"><a href="?page=2" class="page-link">2</a></li>'+

        '  <li class="page-item"><a href="?page=2" class="page-link">&gt;</a></li>'+
        '</ul>'+
        '</nav>',

    props: ['total','current','pp'],

    data: function () {
      return {
          visible: false,
          
      }
    },

    mounted: function () {
        var vm = this;
    },

    methods: {
        onChange: function () {
            console.log(this.visible, this.firstDisabled, this.lastDisabled);
        }
    },

    watch: {
        total: {
            immediate: true,
            handler: function(newVal, oldVal){
                this.onChange();
            }
        }
    }
})
