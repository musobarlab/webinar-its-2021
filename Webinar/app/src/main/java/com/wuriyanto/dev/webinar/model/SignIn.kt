package com.wuriyanto.dev.webinar.model

import com.google.gson.annotations.SerializedName

data class SignIn(
    val username: String,
    val password: String
)

data class SignInResponse(
    @SerializedName("success") val success: Boolean,
    @SerializedName("code") val code: Int,
    @SerializedName("message") val message: String,
    @SerializedName("data") val data: SignInResponseData
)

data class SignInResponseData(

    @SerializedName("accessToken") val accessToken: String,
    @SerializedName("refreshToken") val refreshToken: String,
    @SerializedName("accessTokenExpiresIn") val accessTokenExpiresIn: Int,
    @SerializedName("refreshTokenExpiresIn") val refreshTokenExpiresIn: Int,
    @SerializedName("subscribe") val subscribe: String
)