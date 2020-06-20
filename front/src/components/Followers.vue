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
<!--                    <br/>-->
<!--                    {{ newFollowers}}-->
                </h3>

                <v-row justify="center" v-for="newFollower in newFollowers" :key="newFollower.name">
                    <v-card
                            class="mx-auto"
                            max-width="344"
                            width = "344px"
                    >
                        <v-layout row>
                            <v-img
                                src="http"
                            ></v-img>
                            <v-card-text>
                                <p class="display-1 text--primary">
                                    {{ newFollower.name }}
                                </p>
                                <p>{{ newFollower.description }}</p>
                            </v-card-text>
                        </v-layout>
                        <v-card-actions>
                            <v-btn
                                    text
                                    color="deep-purple accent-4"
                            >
                                +
                            </v-btn>
                            <v-btn
                                    text
                                    color="deep-purple accent-4"
                            >
                                -
                            </v-btn>
                        </v-card-actions>
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

                <v-row justify="center">

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

            axios.get(apiBaseEndpoint + "")
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
            }
        }
    }
</script>
