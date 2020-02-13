Vue.component('pagination', {
    template:
        '<nav>'+
        '<ul class="pagination pagination-sm" v-if="visible">'+
        '  <li class="page-item" v-for="n in pageNumbers()" :class="{active: n==current}">'+
        '    <span class="page-link" v-if="isCurrent(n)">{{n}}</span>'+
        '    <a :href="\'?page=\'+n" class="page-link" v-if="!isCurrent(n)">{{n}}</a>'+
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
        },

        pageNumbers: function(){
            if(this.pageCount<=20){
                return Array.from({length: this.pageCount}, (x,i)=>{return i+1})
            }else{
                var p1 = Array.from({length: 7}, (x,i)=>{return i+1}),
                    p2 = ['...'],
                    p3 = Array.from({length: 7}, (x,i)=>{return this.pageCount-7+i+1});

                if(this.current>=7 && this.current <= this.pageCount-6){

                    if( this.current==7 ){
                        p2 = [8, '...'];
                    }else if( this.current==8 ){
                        p2 = [8, 9, '...'];
                    }else if( this.current==9 ){
                        p2 = [8, 9, 10, '...'];
                    }else if( this.current==this.pageCount-8 ){
                        p2 = ['...', this.pageCount-9, this.pageCount-8, this.pageCount-7];
                    }else if( this.current==this.pageCount-7 ){
                        p2 = ['...', this.pageCount-8, this.pageCount-7];
                    }else if( this.current==this.pageCount-6 ){
                        p2 = ['...', this.pageCount-7];
                    }else{
                        p2 = ['...', this.current-1, this.current, this.current+1, '...'];
                    }

                }

                return p1.concat(p2,p3)
            }

            return [1,2,3]
        },

        isCurrent: function(n){
            return n==this.current || n==='...';
        }
    },

    watch: {
        total: function(){
            this.onChange();
        }
    }

});
