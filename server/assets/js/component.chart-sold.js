Vue.component('chart-sold', {
    template: '<div class="chart-sold"></div>',

    props: ['data'],

    data: function () {
      return {
      }
    },

    mounted: function () {
        var height = 60,
            width = 180,
            marginLeft = 30,
            maxQ = d3.max(this.data, function(d){return d.Q});

        if(maxQ==0){
            return;
        }

        var svg = d3.select(this.$el)
                    .append("svg")
                      .attr("width", width+marginLeft)
                      .attr("height", height+10)
                    .append("g")
                      .attr("transform", "translate("+marginLeft+",5)");

        var x = d3.scaleBand()
                  .range([0, width])
                  .domain(this.data.map(function(d) { return d.D; }))
                  .padding(0.2);

        var xAxis = d3.axisBottom(x)
            .tickValues(x.domain().filter(function(d,i){ return i==0 || i==30 || i==60 || i==90  }))
            .tickFormat((domain,number)=>{return ""});

        svg.append("g")
            .style("font", "8px mono")
            .attr("transform", "translate(0,"+height+")")
            .call(xAxis);

        var y = d3.scaleLinear()
                  .range([height, 0])
                  .domain([0, maxQ]);

        var yAxis = d3.axisLeft(y).ticks(3);

        svg.append("g")
            .style("font", "8px mono")
            .call(yAxis);

        svg.selectAll("mybar")
            .data(this.data)
            .enter()
            .append("rect")
              .attr("x", function(d) { return x(d.D); })
              .attr("y", function(d) { return y(d.Q); })
              .attr("width", function(d) { return x.bandwidth() })
              .attr("height", function(d) { return height - y(d.Q); })
              .style("fill", "#69b3a2")
    },

    methods: {
    },

    watch: {
    }

});
