package com.wuriyanto.dev.webinar

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Button
import android.widget.EditText
import android.widget.ImageView
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.wuriyanto.dev.webinar.model.ProfileResponse
import com.wuriyanto.dev.webinar.model.ProfileUpdateRequest
import com.wuriyanto.dev.webinar.service.ProfileService
import com.wuriyanto.dev.webinar.shared.ServiceBuilder
import com.wuriyanto.dev.webinar.shared.getAccessTokenFromPreference
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class UpdateProfileActivity : AppCompatActivity() {

    private val TAG = "UpdateProfileActivity"

    private lateinit var back: ImageView
    private lateinit var buttonSubmitUpdate: Button
    private lateinit var editTextFullName: EditText

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_update_profile)

        editTextFullName = findViewById(R.id.txt_update_fullname)
        buttonSubmitUpdate = findViewById(R.id.btn_submit_update_profile)
        back = findViewById(R.id.back_from_update_profile)

        back.setOnClickListener {
            startActivity(Intent(this, ProfileActivity::class.java))
            finish()
        }

        buttonSubmitUpdate.setOnClickListener { updateProfile(it) }

        loadProfile()
    }

    private fun updateProfile(view: View) {
        if (editTextFullName.text.toString().isBlank()) {
            showToast("full name cannot be blank")
            return
        }

        val accessToken = getAccessTokenFromPreference(this)
        val profileRequest = ServiceBuilder.buildService(ProfileService::class.java)

        val newFullName = editTextFullName.text.toString()
        val newProfile  = ProfileUpdateRequest(newFullName)
        accessToken?.let {

            profileRequest.updateProfile(it, newProfile).enqueue(
                object : Callback<ProfileResponse> {
                    override fun onFailure(call: Call<ProfileResponse>, t: Throwable) {
                        Log.d(TAG, "update profile fail ${t.message}")
                        onResultUpdate(null)
                    }

                    override fun onResponse(call: Call<ProfileResponse>, response: Response<ProfileResponse>) {
                        Log.d(TAG, "update profile succeed ${response.body().toString()}")
                        onResultUpdate(response.body())
                    }

                }
            )
        }
    }

    private fun onResultUpdate(response: ProfileResponse?) {
        if (response == null) {
            showToast("error getting profile data")
        } else {
            if (response.success) {
                startActivity(Intent(this, ProfileActivity::class.java))
                finish()

            } else {
                showToast("error getting profile data")
            }
        }
    }

    private fun loadProfile() {
        val accessToken = getAccessTokenFromPreference(this)
        val profileRequest = ServiceBuilder.buildService(ProfileService::class.java)
        accessToken?.let {

            profileRequest.getProfile(it).enqueue(
                object : Callback<ProfileResponse> {
                    override fun onFailure(call: Call<ProfileResponse>, t: Throwable) {
                        Log.d(TAG, "get profile fail ${t.message}")
                        onResult(null)
                    }

                    override fun onResponse(call: Call<ProfileResponse>, response: Response<ProfileResponse>) {
                        Log.d(TAG, "get profile succeed ${response.body().toString()}")
                        onResult(response.body())
                    }

                }
            )
        }

    }

    private fun onResult(response: ProfileResponse?) {
        if (response == null) {
            showToast("error getting profile data")
        } else {
            if (response.success) {
                editTextFullName.setText(response.data.fullname)

            } else {
                showToast("error getting profile data")
            }
        }
    }

    private fun showToast(msg: String) {
        Toast.makeText(applicationContext, msg, Toast.LENGTH_SHORT).show()
    }
}