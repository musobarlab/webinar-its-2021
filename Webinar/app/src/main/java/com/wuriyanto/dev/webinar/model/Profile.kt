package com.wuriyanto.dev.webinar.model

import com.google.gson.annotations.SerializedName

data class ProfileResponse(
    @SerializedName("success") val success: Boolean,
    @SerializedName("code") val code: Int,
    @SerializedName("message") val message: String,
    @SerializedName("data") val data: Profile
)

data class Profile (
    @SerializedName("email") val email: String,
    @SerializedName("fullname")val fullname: String,
    @SerializedName("status")val status: String,
    @SerializedName("bonus")val bonus: Float
)

data class ProfileUpdateRequest (
    @SerializedName("fullname")val fullname: String
)