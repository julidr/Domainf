var domainfApp = new Vue({
    el: '#app',
    data() {
        return {
            host: null,
            servers: null,
            serversChanged: null,
            sslGrade: null,
            previousSslGrade: null,
            logo: null,
            title: null,
            isDown: null,
            loading: true,
            loading2: true,
            errored: false,
            hideServers: true,
            gradeColor: null,
            previousColor: null,
            hideHistory: true,
            history: null,
            weirdError: null
        }
    },
    methods: {
        getServers: function () {
            this.loading = true
            this.hideServers = false
            this.hideHistory = true
            var colorBadges = {
                "A+": "badge-success",
                "A": "badge-success",
                "B": "badge-info",
                "C": "badge-warning",
                "D": "badge-warning",
                "F": "badge-danger"
            }
            axios
                .get('http://localhost:8546/servers?host=' + this.host)
                .then(response => {
                    this.servers = response.data.servers
                    this.serversChanged = response.data.servers_changed
                    this.sslGrade = response.data.ssl_grade
                    this.previousSslGrade = response.data.previous_ssl_grade
                    this.logo = response.data.logo
                    this.title = response.data.title
                    this.isDown = response.data.is_down
                    this.gradeColor = colorBadges[this.sslGrade]
                    this.previousColor = colorBadges[this.previousSslGrade]
                })
                .catch(error => {
                    console.log(error)
                    this.weirdError = error
                    this.errored = true
                })
                .finally(() => this.loading = false)
        },
        getHistory: function() {
            this.loading2 = true
            this.hideServers = true
            this.hideHistory = false
            axios
                .get('http://localhost:8546/servers/history')
                .then(response => {
                    this.history = response.data.items
                })
                .catch(error => {
                    console.log(error)
                    this.weirdError = error
                    this.errored = true
                })
                .finally(() => this.loading2 = false)
        }
    }
})