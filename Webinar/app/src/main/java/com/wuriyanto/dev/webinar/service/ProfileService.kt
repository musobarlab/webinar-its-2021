package com.wuriyanto.dev.webinar.service

import com.wuriyanto.dev.webinar.model.ProfileResponse
import com.wuriyanto.dev.webinar.model.ProfileUpdateRequest
import retrofit2.Call
import retrofit2.http.*

interface ProfileService {

    @GET("api/v1/customers/profile")
    fun getProfile(@Header("Authorization") accessToken: String): Call<ProfileResponse>

    @Headers("Content-Type: application/json")
    @PUT("api/v1/customers/profile")
    fun updateProfile(@Header("Authorization") accessToken: String, @Body data: ProfileUpdateRequest): Call<ProfileResponse>
}