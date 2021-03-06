package com.wuriyanto.dev.webinar.service

import com.wuriyanto.dev.webinar.model.SignUp
import com.wuriyanto.dev.webinar.model.SignUpResponse
import retrofit2.Call
import retrofit2.http.Body
import retrofit2.http.Headers
import retrofit2.http.POST

interface SignUpService {

    @Headers("Content-Type: application/json")
    @POST("api/v1/customers")
    fun signUp(@Body data: SignUp): Call<SignUpResponse>
}