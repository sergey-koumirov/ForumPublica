Vue.component('chart-volume', {
    template: '<div class="price-volume"></div>',

    props: ['data'],

    data: function () {
      return {
      }
    },

    mounted: function () {
        this.draw();
    },

    methods: {
        draw: function(){
            if(this.data==null){
                return;
            }

            var height = 60,
                width = 420,
                marginLeft = 60,
                today = new Date(),
                ago90 = new Date();

            today.setHours(23);
            today.setMinutes(59);
            today.setSeconds(59);

            ago90.setDate( today.getDate()-30 )
            ago90.setHours(0);
            ago90.setMinutes(0);
            ago90.setSeconds(0);

            var series = d3.stack()
                    .keys( Array.from(Array( this.data[0][1].length ).keys()) )
                    .value( (d,key)=>{return d[1][key][0]} )(this.data),
                last = series[series.length-1],
                maxVol = d3.max(last, (el)=>{ return el[1]});

            for(var i=0; i<series.length; i++){
                var l = series[i].length;
                for(var j=0; j<l; j++){
                    series[i][j].push( new Date(this.data[j][0]) );
                }
            }

            let colors = this.data[0][1].map( (el)=>{ return el[1]==1 ? '#28a745' : '#f7a588'; } );

            var svg = d3.select(this.$el)
                .append("svg")
                .attr("width", width+marginLeft)
                .attr("height", height+25)
                .append("g")
                .attr("transform", "translate("+marginLeft+",5)");

            var x = d3.scaleUtc()
                    .range([0, width])
                    .domain([ago90, today]),
                xAxis = d3.axisBottom(x)
                    .ticks(d3.utcMonth.every(1));

            svg.append("g")
                .style("font", "8px mono")
                .attr("transform", "translate(0,"+height+")")
                .call(xAxis);

            var y = d3.scaleLinear()
                    .range([height, 0])
                    .domain([0, maxVol]),
                yAxis = d3.axisLeft(y).ticks(3);

            svg.append("g")
                .style("font", "8px mono")
                .call(yAxis);

            var area = d3.area()
                .x( (d) => {return x(d[2]) } )
                .y0( (d) => {return y(d[0]) } )
                .y1( (d) => {return y(d[1]) } );

            svg.append("g")
                .selectAll("path")
                .data(series)
                .join("path")
                .attr("fill", ({index}) => { return colors[index]; })
                .attr("d", area);
        },
    },

    watch: {
    }

});
