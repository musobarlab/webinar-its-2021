package com.wuriyanto.dev.webinar.service

import com.wuriyanto.dev.webinar.model.SignIn
import com.wuriyanto.dev.webinar.model.SignInResponse
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.Headers
import retrofit2.http.POST

interface SignInService {

    @Headers("Content-Type: application/json")
    @POST("api/v1/auth/login")
    fun signIn(@Body data: SignIn): Call<SignInResponse>
}