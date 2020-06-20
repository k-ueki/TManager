<template>
    <v-container>
        <v-row class="text-center">
            <v-col class="mb-4">
                <h4 class="display-2 font-weight-bold mb-3">
                    Followers(New or Bye)
                </h4>
            </v-col>
            <v-tabs centered="true" @change="tabSwitch">
                <v-tab>New</v-tab>
                <v-tab>Bye</v-tab>
            </v-tabs>

            <v-col
                    class="mb-5"
                    cols="12"
                    v-if="tab"
            >
                <h3 class="headline font-weight-bold mb-3">
                    {{ newFollowers.length }} people follow you New!
                </h3>
                <v-btn
                    @click="initFollower"
                >Init</v-btn>

                <v-row justify="center" v-for="newFollower in newFollowers" :key="newFollower.name">
                    <v-card
                            class="mx-auto"
                            max-width="344"
                            width = "344px"
                            v-bind:href="'https://twitter.com/' + newFollower.screen_name"
                    >
                        <v-img
                                style="border-radius: 50%;"
                                height="50%"
                                width="50%"
                                v-bind:src="newFollower.profile_image_url_https"
                        ></v-img>
                        <v-layout row>
                            <v-card-text>
                                <p class="display-1 text--primary">
                                    {{ newFollower.name }}
                                </p>
                                <p>{{ newFollower.description }}</p>
                                <p>Follows : {{newFollower.friends_count}}    Followers : {{newFollower.followers_count}}</p>
                                <p>{{newFollower.description}}</p>
                            </v-card-text>
                        </v-layout>
                    </v-card>
                </v-row>
            </v-col>

            <v-col
                    class="mb-5"
                    cols="12"
                    v-if="!tab"
            >
                <h3 class="headline font-weight-bold mb-3">
                    {{ byeFollowers.length }} people unfollow you!
                </h3>

                <v-btn
                        @click="initFollower"
                >Init</v-btn>

                <v-row justify="center" v-for="byeFollower in byeFollowers" :key="byeFollower.name">
                    <v-card
                            class="mx-auto"
                            max-width="344"
                            width = "344px"
                            v-bind:href="'https://twitter.com/' + byeFollower.screen_name"
                    >
                        <v-img
                                style="border-radius: 50%;"
                                height="50%"
                                width="50%"
                                v-bind:src="byeFollower.profile_image_url_https"
                        ></v-img>
                        <v-layout row>
                            <v-card-text>
                                <p class="display-1 text--primary">
                                    {{ byeFollower.name }}
                                </p>
                                <p>{{ byeFollower.description }}</p>
                                <p>Follows : {{byeFollower.friends_count}}    Followers : {{byeFollower.followers_count}}</p>
                                <p>{{byeFollower.description}}</p>
                            </v-card-text>
                        </v-layout>
                    </v-card>
                </v-row>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    import axios from 'axios'

    const apiBaseEndpoint = "http://localhost:7777/api/v1"


    export default {
        name: 'Followers',
        data (){
            return{
                newFollowers:[],
                byeFollowers:[],
                tab: false,
            }
        },
        created(){
            axios.get(apiBaseEndpoint + "/users/new")
                .then(response => {
                    this.newFollowers = response.data
                })

            axios.get(apiBaseEndpoint + "users/bye")
                .then(response => {
                    this.byeFollowers = response.data
                })
        },
        methods:{
            tabSwitch:function(){
                if(this.tab){
                    this.tab = false
                }else{
                    this.tab = true
                }
            },
            initFollower:function(){
                if(confirm("for sure?")){
                    axios.post(apiBaseEndpoint + "/users/init")
                        .then(resp => {
                            this.newFollower = []
                            console.log(resp)
                        })
                        .catch(err=>{
                            alert("failed to init",err)
                        })
                }
            }
        }
    }
</script>
