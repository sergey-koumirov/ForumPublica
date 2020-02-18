Vue.component('chart-price', {
    template: '<div class="chart-price"></div>',
    props: ['data', 'bottom'],

    data: function () {
      return {
      }
    },

    mounted: function () {

        if(this.data==null){
            return;
        }

        var height = 60,
            width = 180,
            marginLeft = 60,
            maxPrice = d3.max(this.data, function(d){return d.Price}),
            today = new Date(),
            ago90 = new Date();

        today.setHours(23);
        today.setMinutes(59);
        today.setSeconds(59);

        ago90.setDate( today.getDate()-90 )
        ago90.setHours(0);
        ago90.setMinutes(0);
        ago90.setSeconds(0);

        if(maxPrice==0){
            return;
        }

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
            .attr("transform", "translate(0,"+height+")")
            .call(xAxis);

        var y = d3.scaleLinear()
                  .range([height, 0])
                  .domain([0, maxPrice]),
            yAxis = d3.axisLeft(y).ticks(3);

        svg.append("g")
            .call(yAxis);

        var valueline = d3.line()
            .x(function(d) { return x( new Date(d.Dt) ); })
            .y(function(d) { return y( d.Price ); });

        svg.append("path")
            .data([this.data])
            .attr("class", "line")
            .attr("d", valueline)
            .style("fill", "none")
            .style("stroke", "#69b3a2");

        if(!!this.bottom && this.bottom>0){
            svg.append("path")
                .datum([{Dt: ago90, Price: this.bottom},{Dt: today, Price: this.bottom}])
                .attr("class", "line")
                .attr("d", valueline)
                .style("fill", "none")
                .style("stroke", "#b369a2")
                .style("stroke-dasharray", "2 4");
        }

    },

    methods: {
    },

    watch: {
    }

});
