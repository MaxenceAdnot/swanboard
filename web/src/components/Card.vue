<template>
  <div class="col-md-4">
      <div class="card card-profile" v-bind:class="{ offline: !status }">
          <div class="card-avatar">
              <a :href="'http://' + ip"> <img class="img" :src="avatar"> </a>
          </div>
          <div class="table">
              <h4 class="card-caption">{{ username }}</h4>
              <h6 v-bind:class="{ 'text-muted': status, 'text-light': !status }">{{ ip }}</h6>
              <div v-if="status">
                <div v-if="portsList.length > 0" class="d-flex justify-content-around flex-row">
                  <p v-for="port in portsList" :key="port" class="card-description"><span class="circle"></span>{{ port }}</p>
                </div>
                <div v-else>
                  <p class="card-description">No open ports</p>
                </div>
              </div>
              <div v-else>
                <p class="card-description">offline</p>
              </div>
          </div>
      </div>
  </div>
</template>

<script>
const Trianglify = require('trianglify')

export default {
  name: 'Card',
  props: {
    username: String,
    ip: String,
    ports: String,
    status: Boolean
  },
  data () {
    return {
      portsList: [],
      avatar: 'data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAPoAAAD6AQMAAACyIsh+AAAABlBMVEX///////9VfPVsAAAACXBIWXMAAA7EAAAOxAGVKw4bAAAAxUlEQVRoge2RQQ4CIQxFGYeYbkyaeAF2HkOONkfxKHMslxZQjKG4Nua9Bfmhn9+mhAAAAAAAAADw36zl2LrKgyElO7QrGSO0XxcV3TbalRNQeb87uPXkZH2UzZBiLU8MFqzS3i/OhJfNDGJFPe+eoSXENE1IdfZDbsox5GpYtqa8FradZij7mBhCbTFbdBBrXocM0d1DNYh+MZRofX7T+BdlQVa1VZsSGaew+CQvpc6U1/tpW3NTx325TYYAAAAAAAAA+BUerdwNjuTEMqIAAAAASUVORK5CYII='
    }
  },
  mounted () {
    if (this.ports !== '') {
      this.portsList = this.ports.split(',')
    }

    var pattern = Trianglify({
      seed: this.username + 'swanboardipsec',
      cell_size: 50,
      width: 300,
      height: 300
    })
    this.avatar = pattern.png()
  }
}
</script>

<style scoped>
.card {
    display: inline-block;
    position: relative;
    width: 100%;
    margin-bottom: 30px;
    border-radius: 6px;
    color: rgba(0, 0, 0, 0.87);
    background: #fff;
    box-shadow: 0 2px 2px 0 rgba(0, 0, 0, 0.14), 0 3px 1px -2px rgba(0, 0, 0, 0.2), 0 1px 5px 0 rgba(0, 0, 0, 0.12);
}

.card img {
    width: 100%;
    height: auto;
}

.offline {
    color: #ffffff;
    background: #c0b6b69f;
}

.card-profile {
    margin-top: 50px;
    text-align: center;
}

.card-profile .card-avatar {
    max-width: 130px;
    max-height: 130px;
    margin: -50px auto 0;
    border-radius: 50%;
    overflow: hidden;
    box-shadow: 0 16px 38px -12px rgba(0, 0, 0, 0.56), 0 4px 25px 0px rgba(0, 0, 0, 0.12), 0 8px 10px -5px rgba(0, 0, 0, 0.2);
}

.card .table {
    margin-bottom: 0px;
    padding: 15px 30px;
}

.circle:before {
  content: ' \25CF';
  color: green;
  font-size: 18px;
  margin: 0px 5px;
}
</style>
