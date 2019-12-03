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
            errored: false,
            hideServers: true
        }
    },
    methods: {
        getServers: function () {
            this.loading = true
            this.hideServers = false
            axios
                .get('http://localhost:8546/servers?host=' + this.host)
                .then(response => {
                    this.servers = response.data.servers,
                        this.serversChanged = response.data.servers_changed,
                        this.sslGrade = response.data.ssl_grade,
                        this.previousSslGrade = response.data.previous_ssl_grade,
                        this.logo = response.data.logo,
                        this.title = response.data.title,
                        this.isDown = response.data.is_down
                })
                .catch(error => {
                    console.log(error)
                    this.errored = true
                })
                .finally(() => this.loading = false)
        }
    }
})