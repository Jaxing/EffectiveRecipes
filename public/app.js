new Vue({
    el: '#app',

    data: {
        ws: null, // Our websocket,
        query: "",
        recipes: [],
        activeRecipe: "Recipe name"
    },
    created: function() {
        var self = this;
        this.ws = new WebSocket('ws://' + window.location.host + '/ws');
        this.ws.addEventListener('message', function(e) {
            var message = JSON.parse(e.data);
            var recipes = message.recipes;
            for (i = 0; i < recipes.length; i++) {
                var recipe = recipes[i];
                self.recipes.push(recipe)
            }
            self.activeRecipe = recipes[0]
        });
    },
    methods: {
        search: function () {
            if (this.query != '') {
                var msg = JSON.stringify({
                    query: this.query,
                });
                console.debug("Message sent: " + msg)
                this.ws.send(msg);
                this.query = ''; // Reset newMsg
            }
        },
        showRecipe: function (recipe) {
            this.activeRecipe = recipe;
        },
        //this.email = $('<p>').html(this.email).text();
    }
});