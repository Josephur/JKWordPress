package config

/*
 * The base configuration for WordPress (originally wp-config.php) re-written for Go
 *
 * The jkwp-config.php creation script uses this file during the installation.
 * You don't have to use the website; you can copy and modify this file to fill
 * in the values.
 *
 * This file contains the following configurations:
 *
 *  * Database settings
 *  * Secret keys
 *  * Database table prefix
 *  * Debug
 *
 * @link https://developer.wordpress.org/advanced-administration/wordpress/wp-config/
 *
 * @package WordPress
 */

// Config holds all the configuration values
type Config struct {
	// ** Database settings - You can get this info from your web host **
	// The name of the database for WordPress
	DBName string

	// Database username
	DBUser string

	// Database password
	DBPassword string

	// Database hostname
	DBHost string

	// Database charset to use in creating database tables
	DBCharset string

	// The database collate type. Don't change this if in doubt.
	DBCollate string

	/*
	 * Authentication unique keys and salts.
	 *
	 * Change these to different unique phrases! You can generate these using
	 * the https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service.
	 *
	 * You can change these at any point in time to invalidate all existing cookies.
	 * This will force all users to have to log in again.
	 */
	AuthKey        string
	SecureAuthKey  string
	LoggedInKey    string
	NonceKey       string
	AuthSalt       string
	SecureAuthSalt string
	LoggedInSalt   string
	NonceSalt      string

	/*
	 * WordPress database table prefix.
	 *
	 * You can have multiple installations in one database if you give each
	 * a unique prefix. Only numbers, letters, and underscores please!
	 *
	 * At the installation time, database tables are created with the specified prefix.
	 * Changing this value after WordPress is installed will make your site think
	 * it has not been installed.
	 */
	TablePrefix string

	/*
	 * For developers: WordPress debugging mode.
	 *
	 * Change this to true to enable the display of notices during development.
	 * It is strongly recommended that plugin and theme developers use WP_DEBUG
	 * in their development environments.
	 */
	WPDebug bool
}

// LoadConfig returns a Config with default WordPress-like settings.
// You can customize this function to load from environment variables,
// a JSON/YAML file, or other sources as desired.
func LoadConfig() *Config {
	return &Config{
		// ** Database settings **
		DBName:     "database_name_here",
		DBUser:     "username_here",
		DBPassword: "password_here",
		DBHost:     "localhost",
		DBCharset:  "utf8",
		DBCollate:  "",

		// ** Authentication unique keys and salts **
		AuthKey:        "put your unique phrase here",
		SecureAuthKey:  "put your unique phrase here",
		LoggedInKey:    "put your unique phrase here",
		NonceKey:       "put your unique phrase here",
		AuthSalt:       "put your unique phrase here",
		SecureAuthSalt: "put your unique phrase here",
		LoggedInSalt:   "put your unique phrase here",
		NonceSalt:      "put your unique phrase here",

		// ** Table prefix **
		TablePrefix: "jkwp_",

		// ** Debug mode **
		WPDebug: false,
	}
}
