<template>
  <header class="fullscreen d-flex flex-column">
    <canvas id="header"></canvas>
    <section class="container text-center my-auto">
      <h1 class="mb-1">Swanboard</h1>
      <h3 class="mb-5">A simple ipsec wrapper</h3>
    </section>
    <section class="container">
        <div class="row">
          <Card v-for="user in usersInfo" :key="user.ip" :username="user.username" :ports="user.ports" :ip="user.ip" :status="user.status === 'online' ? true : false"/>
        </div>
      </section>
      <resize-observer @notify="handleResize" />
  </header>
</template>

<script>
import Card from './components/Card.vue'

const axios = require('axios')
const Trianglify = require('trianglify')

export default {
  name: 'app',
  components: {
    Card
  },
  data () {
    return {
      usersInfo: [],
      loading: true,
      errored: false,
      error: null,
      currentHeight: 0
    }
  },
  mounted () {
    this.currentHeight = document.documentElement.scrollHeight
    var pattern = Trianglify({
      x_colors: 'YlGnBu',
      cell_size: 50,
      width: document.documentElement.scrollWidth,
      height: this.currentHeight
    })
    var head = document.getElementById('header')
    pattern.canvas(head)

    axios
      .get('/api/getusers')
      .then(response => {
        this.usersInfo = response.data
      })
      .catch(error => {
        this.error = error
        this.errored = true
      })
      .finally(() => {
        this.loading = false
      })
  },
  methods: {
    handleResize () {
      if (this.currentHeight !== document.documentElement.scrollHeight) {
        this.currentHeight = document.documentElement.scrollHeight
        var pattern = Trianglify({
          x_colors: 'YlGnBu',
          width: document.documentElement.scrollWidth,
          height: this.currentHeight
        })
        var head = document.getElementById('header')
        pattern.canvas(head)
      }
    }
  }
}
</script>

<style scoped>
body,
html {
    font-family: 'Open Sans', sans-serif;
}

.fullscreen {
    position: relative;
    width: 100%;
    height: 100%;
}

#header {
    position: absolute;
    left: 50%;
    top: 50%;
    -webkit-transform: translate(-50%,-50%);
    transform: translate(-50%,-50%);
    min-width: 100%;
    z-index: -1;
}

.fullscreen h1 {
    font-size: 4rem;
    font-family: 'Courgette', cursive;
    margin: 0;
    padding: 0;
}

.fullscreen h1,h3{
    color: #ffffff;
}

@media (min-width: 768px) {
    .fullscreen {
        min-height: 100vh;
    }
    .fullscreen h1 {
        font-size: 5.5rem;
    }
}

@media (max-width: 768px) {
    .fullscreen {
        padding-top: 25px;
    }
}
</style>
