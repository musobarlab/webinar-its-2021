package com.wuriyanto.dev.webinar

import android.content.Intent
import android.os.Bundle
import android.util.Log
import android.view.View
import android.widget.Button
import android.widget.TextView
import android.widget.Toast
import androidx.appcompat.app.AppCompatActivity
import com.wuriyanto.dev.webinar.model.Profile
import com.wuriyanto.dev.webinar.model.ProfileResponse
import com.wuriyanto.dev.webinar.service.ProfileService
import com.wuriyanto.dev.webinar.shared.ServiceBuilder
import com.wuriyanto.dev.webinar.shared.getAccessTokenFromPreference
import com.wuriyanto.dev.webinar.shared.removeAccessTokenToPreference
import com.wuriyanto.dev.webinar.shared.setAccessTokenToPreference
import retrofit2.Call
import retrofit2.Callback
import retrofit2.Response

class ProfileActivity : AppCompatActivity() {

    private val TAG  = "ProfileActivity"

    private lateinit var buttonLogout: Button
    private lateinit var buttonPrepareEdit: Button

    private lateinit var txtEmail: TextView
    private lateinit var txtName: TextView
    private lateinit var txtStatus: TextView
    private lateinit var txtBonus: TextView

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_profile)

        txtEmail = findViewById(R.id.txt_profile_email)
        txtName = findViewById(R.id.txt_profile_name)
        txtStatus = findViewById(R.id.txt_profile_status)
        txtBonus = findViewById(R.id.txt_profile_bonus)

        buttonLogout = findViewById(R.id.btn_profile_logout)
        buttonLogout.setOnClickListener { logout(it) }

        buttonPrepareEdit = findViewById(R.id.btn_prepare_update)
        buttonPrepareEdit.setOnClickListener { prepareEdit(it) }

        loadProfile()
    }

    private fun logout(v: View) {
        removeAccessTokenToPreference(this)
        startActivity(Intent(this, SignInActivity::class.java))
        finish()
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
                Log.d(TAG, "email ${response.data.email}")
                txtEmail.text = response.data.email
                txtName.text = response.data.fullname
                txtStatus.text = response.data.status
                txtBonus.text = "Bonus : ${response.data.bonus}"

            } else {
                showToast("error getting profile data")
            }
        }
    }

    private fun showToast(msg: String) {
        Toast.makeText(applicationContext, msg, Toast.LENGTH_SHORT).show()
    }

    private fun prepareEdit(view: View) {
        startActivity(Intent(this, UpdateProfileActivity::class.java))
        finish()
    }
}