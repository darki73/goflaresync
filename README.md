# GoFlareSync
Simple application to sync your Cloudflare DNS records with your local IP address.

# Installation
This section is under construction.  
As of now, you can download the latest release from the releases page.

# Configuration
This section describes how to configure the application.  
The configuration file is located at `/etc/goflaresync/config.yaml` by default.

## Credentials
While configuring the application, you will be asked to provide your Cloudflare credentials.  
These credentials are used to authenticate with the Cloudflare API.  

Here is an example of the credentials section:
```yaml
credentials:
  email: administrator@example.com
  token: 1234567890qwerty
```

## Records
While configuring the application, you will be asked to provide your Cloudflare DNS records.  
These records are used to identify the DNS records that should be monitored and updated in the case of an IP address change.  

Here is an example of the records section:
```yaml
records:
  - name: example.com
    type: A
    proxied: true
  - name: subdomain.example.com
    type: A
```

It is worth noting that all of the existing attributes of the record will be preserved upon the update.  
The only attribute that will be changed is the IP address (content).

## Watcher
This section describes the watcher configuration and is optional.  
By default, the watcher will check for an IP address change every 5 minutes and will use `https://api.ipify.org` to retrieve the IP address.

If you want to change the interval or the IP address provider, you can do so by adding the following section to your configuration file:
```yaml
watcher:
  interval: 5m
  address_source: https://api.ipify.org
```

## Log Level
This section describes the log level configuration and is optional.

By default, the log level is set to `info`.  
If you want to change the log level, you can do so by adding the following section to your configuration file:
```yaml
log_level: debug
```

**Supported log levels:**
* `t` | `trace` | `Trace` | `TRACE` - displays all log messages
* `d` | `debug` | `Debug` | `DEBUG` - displays debug log messages
* `i` | `info` | `Info` | `INFO` - displays info log messages
* `w` | `warn` | `warning` | `Warn` | `Warning` | `WARN` | `WARNING` - displays warn log messages
* `e` | `err` | `error` | `Err` | `Error` | `ERR` | `ERROR` - displays error log messages
* `f` | `fatal` | `Fatal` | `FATAL` - displays fatal log messages
* `p` | `panic` | `Panic` | `PANIC` - displays panic log messages

## Complete Configuration Example
Here is a complete configuration example:
```yaml
credentials:
  email: administrator@example.com
  token: 1234567890qwerty
records:
  - name: example.com
    type: A
    proxied: true
  - name: subdomain.example.com
    type: A
watcher:
  interval: 5m
  address_source: https://api.ipify.org
log_level: debug
```

# Usage
This section describes how to use the application.  

## Flags
The following flags are available:
* `--configuration-path` - Path to the configuration file. (default: `/etc/goflaresync`)
* `--configuration-name` - Name of the configuration file. (default: `config`)
* `--configuration-extension` - Extension of the configuration file. (default: `yaml`)

## Commands
The following commands are available:
* `help` - Help about any command
* `start` - Start the application
* `version` - Print the version number of GoFlareSync
* `configuration` - Meta command that provides access to configuration related commands
  * `configuration display` - Displays the current configuration (omits the E-Mail and API Key)
  * `configuration display-full` - Displays the current configuration (includes the E-Mail and API Key)
* `service` - Meta command that provides access to service related commands
  * `service install` - Installs the application as a service (also enables the service to start on boot)
  * `service uninstall` - Uninstalls the application as a service
  * `service start` - Starts the application as a service
  * `service stop` - Stops the application as a service
  * `service restart` - Restarts the application as a service
  * `service enable` - Enables the service to start on boot
  * `service disable` - Disables the service to start on boot