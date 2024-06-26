---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_ec2_security_group_egress Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Adds the specified outbound (egress) rule to a security group.
   An outbound rule permits instances to send traffic to the specified IPv4 or IPv6 address range, the IP addresses that are specified by a prefix list, or the instances that are associated with a destination security group. For more information, see Security group rules https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html.
   You must specify exactly one of the following destinations: an IPv4 or IPv6 address range, a prefix list, or a security group. Otherwise, the stack launches successfully but the rule is not added to the security group.
   You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also specify the ICMP/ICMPv6 type and code. To specify all types or all codes, use -1.
   Rule changes are propagated to instances associated with the security group as quickly as possible
---

# awscc_ec2_security_group_egress (Resource)

Adds the specified outbound (egress) rule to a security group.
 An outbound rule permits instances to send traffic to the specified IPv4 or IPv6 address range, the IP addresses that are specified by a prefix list, or the instances that are associated with a destination security group. For more information, see [Security group rules](https://docs.aws.amazon.com/vpc/latest/userguide/security-group-rules.html).
 You must specify exactly one of the following destinations: an IPv4 or IPv6 address range, a prefix list, or a security group. Otherwise, the stack launches successfully but the rule is not added to the security group.
 You must specify a protocol for each rule (for example, TCP). If the protocol is TCP or UDP, you must also specify a port or port range. If the protocol is ICMP or ICMPv6, you must also specify the ICMP/ICMPv6 type and code. To specify all types or all codes, use -1.
 Rule changes are propagated to instances associated with the security group as quickly as possible



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `group_id` (String) The ID of the security group. You must specify either the security group ID or the security group name in the request. For security groups in a nondefault VPC, you must specify the security group ID.
- `ip_protocol` (String) The IP protocol name (``tcp``, ``udp``, ``icmp``, ``icmpv6``) or number (see [Protocol Numbers](https://docs.aws.amazon.com/http://www.iana.org/assignments/protocol-numbers/protocol-numbers.xhtml)).
 Use ``-1`` to specify all protocols. When authorizing security group rules, specifying ``-1`` or a protocol number other than ``tcp``, ``udp``, ``icmp``, or ``icmpv6`` allows traffic on all ports, regardless of any port range you specify. For ``tcp``, ``udp``, and ``icmp``, you must specify a port range. For ``icmpv6``, the port range is optional; if you omit the port range, traffic for all types and codes is allowed.

### Optional

- `cidr_ip` (String) The IPv4 address range, in CIDR format.
 You must specify a destination security group (``DestinationPrefixListId`` or ``DestinationSecurityGroupId``) or a CIDR range (``CidrIp`` or ``CidrIpv6``).
 For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
- `cidr_ipv_6` (String) The IPv6 address range, in CIDR format.
 You must specify a destination security group (``DestinationPrefixListId`` or ``DestinationSecurityGroupId``) or a CIDR range (``CidrIp`` or ``CidrIpv6``).
 For examples of rules that you can add to security groups for specific access scenarios, see [Security group rules for different use cases](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/security-group-rules-reference.html) in the *User Guide*.
- `description` (String) The description of an egress (outbound) security group rule.
 Constraints: Up to 255 characters in length. Allowed characters are a-z, A-Z, 0-9, spaces, and ._-:/()#,@[]+=;{}!$*
- `destination_prefix_list_id` (String) The prefix list IDs for an AWS service. This is the AWS service that you want to access through a VPC endpoint from instances associated with the security group.
 You must specify a destination security group (``DestinationPrefixListId`` or ``DestinationSecurityGroupId``) or a CIDR range (``CidrIp`` or ``CidrIpv6``).
- `destination_security_group_id` (String) The ID of the security group.
 You must specify a destination security group (``DestinationPrefixListId`` or ``DestinationSecurityGroupId``) or a CIDR range (``CidrIp`` or ``CidrIpv6``).
- `from_port` (Number) If the protocol is TCP or UDP, this is the start of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP type or -1 (all ICMP types).
- `to_port` (Number) If the protocol is TCP or UDP, this is the end of the port range. If the protocol is ICMP or ICMPv6, this is the ICMP code or -1 (all ICMP codes). If the start port is -1 (all ICMP types), then the end port must be -1 (all ICMP codes).

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `security_group_egress_id` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_ec2_security_group_egress.example <resource ID>
```
