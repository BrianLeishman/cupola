package main

import "fmt"

// reference types
const (
	pstInt16                uint16 = 0x0002
	pstInt32                       = 0x0003
	pstFloat32                     = 0x0004
	pstFloat64                     = 0x0005
	pstInt64                       = 0x0006
	pstApplicationTime             = 0x0007
	pst32BitErrorValue             = 0x000A
	pstBool                        = 0x000B
	pstEmbeddedObject              = 0x000D
	_pstInt64                      = 0x0014
	pstNullTerminatedString        = 0x001E
	pstString                      = 0x001F
	pstSystime                     = 0x0040
	pstOLEGuid                     = 0x0048
	pstBytes                       = 0x0102
	pstInt32Array                  = 0x1003
	pstInt64Array                  = 0x1014
	pstStringArray                 = 0x101E
	pstBytesArray                  = 0x1102
)

func referenceName(n uint16) string {
	switch n {
	case pstInt16:
		return "Signed 16bit value"
	case pstInt32:
		return "Signed 32bit value"
	case pstFloat32:
		return "4-byte floating point"
	case pstFloat64:
		return "Floating point double"
	case pstInt64:
		return "Signed 64-bit int"
	case pstApplicationTime:
		return "Application Time"
	case pst32BitErrorValue:
		return "32-bit error value"
	case pstBool:
		return "Boolean (non-zero = true)"
	case pstEmbeddedObject:
		return "Embedded Object"
	case _pstInt64:
		return "8-byte signed integer (64-bit)"
	case pstNullTerminatedString:
		return "Null terminated String"
	case pstString:
		return "Unicode string"
	case pstSystime:
		return "Systime - Filetime structure"
	case pstOLEGuid:
		return "OLE Guid"
	case pstBytes:
		return "Binary data"
	case pstInt32Array:
		return "Array of 32bit values"
	case pstInt64Array:
		return "Array of 64bit values"
	case pstStringArray:
		return "Array of Strings"
	case pstBytesArray:
		return "Array of Binary data"
	}
	return fmt.Sprintf("0x%04x", n)
}

// item types
const (
	pstItemAlternateRecipientAllowed = 0x0002 // Alternate recipient allowed
	pstItemExtendedAttributesTable = 0x0003 // Extended Attributes Table
	pstItemImportanceLevel = 0x0017 // Importance Level
	pstItemIPMContextMessageClass = 0x001a // IPM Context, message class
	pstItemGlobalDeliveryReportRequested = 0x0023 // Global delivery report requested
	pstItemPriority = 0x0026 // Priority
	pstItemReadReceipt = 0x0029 // Read Receipt
	pstItemReassignmentProhibited = 0x002b // Reassignment Prohibited
	pstItemOriginalSensitivity = 0x002e // Original Sensitivity
	pstItemReportTime = 0x0032 // Report time
	pstItemSensitivity = 0x0036 // Sensitivity
	pstItemEmailSubject = 0x0037 // Email Subject
	pstItemClientSubmitTimeDateSent = 0x0039 // Client submit time / date sent
	pstItemOutlookAddressOfSender = 0x003b // Outlook Address of Sender
	pstItemOutlookStructureDescribingTheRecipient = 0x003f // Outlook structure describing the recipient
	pstItemNameOfTheOutlookRecipientStructure = 0x0040 // Name of the Outlook recipient structure
	pstItemOutlookStructureDescribingTheSender = 0x0041 // Outlook structure describing the sender
	pstItemNameOfTheOutlookSenderStructure = 0x0042 // Name of the Outlook sender structure
	pstItemAnotherStructureDescribingTheRecipient = 0x0043 // Another structure describing the recipient
	pstItemNameOfTheSecondRecipientStructure = 0x0044 // Name of the second recipient structure
	pstItemReplyToOutlookStructure = 0x004f // Reply-To Outlook Structure
	pstItemNameOfTheReplyToStructure = 0x0050 // Name of the Reply-To structure
	pstItemOutlookNameOfRecipient = 0x0051 // Outlook Name of recipient
	pstItemSecondOutlookNameOfRecipient = 0x0052 // Second Outlook name of recipient
	pstItemMyAddressInTOField = 0x0057 // My address in TO field
	pstItemMyAddressInCCField = 0x0058 // My address in CC field
	pstItemMessageAddressedToMe = 0x0059 // Message addressed to me
	pstItemResponseRequested = 0x0063 // Response requested
	pstItemSendersAddressAccessMethodSMTPEX = 0x0064 // Sender's Address access method (SMTP, EX)
	pstItemSendersAddress = 0x0065 // Sender's Address
	pstItemConversationTopicProcessedSubjectwithFwdReRemoved = 0x0070 // Conversation topic, processed subject (with Fwd:, Re, ... removed)
	pstItemConversationIndex = 0x0071 // Conversation index
	pstItemOriginalDisplayBCC = 0x0072 // Original display BCC
	pstItemOriginalDisplayCC = 0x0073 // Original display CC
	pstItemOriginalDisplayTO = 0x0074 // Original display TO
	pstItemRecipientAddressAccessMethodSMTPEX = 0x0075 // Recipient Address Access Method (SMTP, EX)
	pstItemRecipientsAddress = 0x0076 // Recipient's Address
	pstItemSecondRecipientAccessMethodSMTPEX = 0x0077 // Second Recipient Access Method (SMTP, EX)
	pstItemSecondRecipientAddress = 0x0078 // Second Recipient Address
	pstItemEmailHeader = 0x007d // Email Header. This is the header that was attached to the email
	pstItemNDRReasonCode = 0x0c04 // NDR Reason code
	pstItemNDRDiagCode = 0x0c05 // NDR Diag code
	pstItemNonreceiptNotificationRequested = 0x0c06 // Non-receipt notification requested
	pstItemReplyRequested = 0x0c17 // Reply Requested
	pstItemSecondSenderStructure = 0x0c19 // Second sender structure
	pstItemNameOfSecondSenderStructure = 0x0c1a // Name of second sender structure
	pstItemSupplementaryInfo = 0x0c1b // Supplementary info
	pstItemSecondOutlookNameOfSender = 0x0c1d // Second outlook name of sender
	pstItemSecondSenderAccessMethodSMTPEX = 0x0c1e // Second sender access method (SMTP, EX)
	pstItemSecondSenderAddress = 0x0c1f // Second Sender Address
	pstItemNDRStatusCode = 0x0c20 // NDR status code
	pstItemDeleteAfterSubmit = 0x0e01 // Delete after submit
	pstItemBCCAddresses = 0x0e02 // BCC Addresses
	pstItemCCAddresses = 0x0e03 // CC Addresses
	pstItemSentToAddress = 0x0e04 // SentTo Address
	pstItemDate = 0x0e06 // Date.
	pstItemFlagBits = 0x0e07 // Flag bits
	pstItemMessageSize = 0x0e08 // Message Size
	pstItemSentmailEntryID = 0x0e0a // Sentmail EntryID
	pstItemCompressedRTFInSync = 0x0e1f // Compressed RTF in Sync
	pstItemAttachmentSize = 0x0e20 // Attachment Size
	pstItembinaryRecordHeader = 0x0ff9 // binary record header
	pstItemPlainTextEmailBody = 0x1000 // Plain Text Email Body. Does not exist if the email doesn't have a plain text version
	pstItemReportText = 0x1001 // Report Text
	pstItemRTFSyncBodyCRC = 0x1006 // RTF Sync Body CRC
	pstItemRTFSyncBodyCharacterCount = 0x1007 // RTF Sync Body character count
	pstItemRTFSyncBodyTag = 0x1008 // RTF Sync body tag
	pstItemRTFCompressedBody = 0x1009 // RTF Compressed body
	pstItemRTFWhitespacePrefixCount = 0x1010 // RTF whitespace prefix count
	pstItemRTFWhitespaceTailingCount = 0x1011 // RTF whitespace tailing count
	pstItemHTMLEmailBody = 0x1013 // HTML Email Body. Does not exist if the email doesn't have an HTML version
	pstItemMessageID = 0x1035 // Message ID
	pstItemInReplyToOrParentsMessageID = 0x1042 // In-Reply-To or Parent's Message ID
	pstItemReturnPath = 0x1046 // Return Path
	pstItemFolderName = 0x3001 // Folder Name? I have seen this value used for the contacts record aswell
	pstItemAddressType = 0x3002 // Address Type
	pstItemContactAddress = 0x3003 // Contact Address
	pstItemComment = 0x3004 // Comment
	pstItemDateItemCreation = 0x3007 // Date item creation
	pstItemDateItemModification = 0x3008 // Date item modification
	pstItembinaryRecordHeader2 = 0x300b // binary record header
	pstItemValidFolderMask = 0x35df // Valid Folder Mask
	pstItemBinaryRecordContainsAReferenceToTopOfPersonalFolderItem = 0x35e0 // binary record contains a reference to "Top of Personal Folder" item
	pstItemBinaryRecordContainsAReferenceToDefaultOutboxItem = 0x35e2 // binary record contains a reference to default outbox item
	pstItemBinaryRecordContainsAReferenceToDeletedItemsItem = 0x35e3 // binary record contains a reference to "Deleted Items" item
	pstItemBinaryRecordContainsAReferenceToSentItemsFolderItem = 0x35e4 // binary record contains a reference to sent items folder item
	pstItemBinaryRecordContainsAReferenceToUserViewsFolderItem = 0x35e5 // binary record contains a reference to user views folder item
	pstItemBinaryRecordContainsAReferenceToCommonViewsFolderItem = 0x35e6 // binary record contains a reference to common views folder item
	pstItemBinaryRecordContainsAReferenceToSearchRootItem = 0x35e7 // binary record contains a reference to "Search Root" item
	pstItemNumberOfEmailsStoredInAFolder = 0x3602 // the number of emails stored in a folder
	pstItemNumberOfUnreadEmailsInAFolder = 0x3603 // the number of unread emails in a folder
	pstItemHasSubfolders = 0x360a // Has Subfolders
	pstItemFolderContentDescription = 0x3613 // the folder content description
	pstItemAssociateContentCount = 0x3617 // Associate Content count
	pstItemBinaryDataAttachment = 0x3701 // Binary Data attachment
	pstItemAttachmentFilename = 0x3704 // Attachment Filename
	pstItemAttachmentMethod = 0x3705 // Attachement method
	pstItemAttachmentFilenameLong = 0x3707 // Attachment Filename long
	pstItemAttachmentPosition = 0x370b // Attachment Position
	pstItemAttachmentMimeEncoding = 0x370e // Attachment mime encoding
	pstItemAttachmentMimeSequence = 0x3710 // Attachment mime Sequence
	pstItemContactsAccountName = 0x3a00 // Contact's Account name
	pstItemContactAlternateRecipient = 0x3a01 // Contact Alternate Recipient
	pstItemCallbackTelephoneNumber = 0x3a02 // Callback telephone number
	pstItemMessageConversionProhibited = 0x3a03 // Message Conversion Prohibited
	pstItemContactsSuffix = 0x3a05 // Contacts Suffix
	pstItemContactsFirstName = 0x3a06 // Contacts First Name
	pstItemContactsGovernmentIDNumber = 0x3a07 // Contacts Government ID Number
	pstItemBusinessTelephoneNumber = 0x3a08 // Business Telephone Number
	pstItemHomeTelephoneNumber = 0x3a09 // Home Telephone Number
	pstItemContactsInitials = 0x3a0a // Contacts Initials
	pstItemKeyword = 0x3a0b // Keyword
	pstItemContactsLanguage = 0x3a0c // Contact's Language
	pstItemContactsLocation = 0x3a0d // Contact's Location
	pstItemMailPermission = 0x3a0e // Mail Permission
	pstItemMHSCommonName = 0x3a0f // MHS Common Name
	pstItemOrganizationalID = 0x3a10 // Organizational ID #
	pstItemContactsSurname = 0x3a11 // Contacts Surname
	pstItemoriginalEntryID = 0x3a12 // original entry id
	pstItemoriginalDisplayName = 0x3a13 // original display name
	pstItemoriginalSearchKey = 0x3a14 // original search key
	pstItemDefaultPostalAddress = 0x3a15 // Default Postal Address
	pstItemCompanyName = 0x3a16 // Company Name
	pstItemJobTitle = 0x3a17 // Job Title
	pstItemDepartmentName = 0x3a18 // Department Name
	pstItemOfficeLocation = 0x3a19 // Office Location
	pstItemPrimaryTelephone = 0x3a1a // Primary Telephone
	pstItemBusinessPhoneNumber2 = 0x3a1b // Business Phone Number 2
	pstItemMobilePhoneNumber = 0x3a1c // Mobile Phone Number
	pstItemRadioPhoneNumber = 0x3a1d // Radio Phone Number
	pstItemCarPhoneNumber = 0x3a1e // Car Phone Number
	pstItemOtherPhoneNumber = 0x3a1f // Other Phone Number
	pstItemTransmittableDisplayName = 0x3a20 // Transmittable Display Name
	pstItemPagerPhoneNumber = 0x3a21 // Pager Phone Number
	pstItemuserCertificate = 0x3a22 // user certificate
	pstItemPrimaryFaxNumber = 0x3a23 // Primary Fax Number
	pstItemBusinessFaxNumber = 0x3a24 // Business Fax Number
	pstItemHomeFaxNumber = 0x3a25 // Home Fax Number
	pstItemBusinessAddressCountry = 0x3a26 // Business Address Country
	pstItemBusinessAddressCity = 0x3a27 // Business Address City
	pstItemBusinessAddressState = 0x3a28 // Business Address State
	pstItemBusinessAddressStreet = 0x3a29 // Business Address Street
	pstItemBusinessPostalCode = 0x3a2a // Business Postal Code
	pstItemBusinessPOBox = 0x3a2b // Business PO Box
	pstItemTelexNumber = 0x3a2c // Telex Number
	pstItemISDNNumber = 0x3a2d // ISDN Number
	pstItemAssistantPhoneNumber = 0x3a2e // Assistant Phone Number
	pstItemHomePhone2 = 0x3a2f // Home Phone 2
	pstItemAssistantsName = 0x3a30 // Assistant's Name
	pstItemCanReceiveRichText = 0x3a40 // Can receive Rich Text
	pstItemWeddingAnniversary = 0x3a41 // Wedding Anniversary
	pstItemBirthday = 0x3a42 // Birthday
	pstItemHobbies = 0x3a43 // Hobbies
	pstItemMiddleName = 0x3a44 // Middle Name
	pstItemDisplayNamePrefixTitle = 0x3a45 // Display Name Prefix (Title)
	pstItemProfession = 0x3a46 // Profession
	pstItemPreferredByName = 0x3a47 // Preferred By Name
	pstItemSpousesName = 0x3a48 // Spouse's Name
	pstItemComputerNetworkName = 0x3a49 // Computer Network Name
	pstItemCustomerID = 0x3a4a // Customer ID
	pstItemTTYTDDPhone = 0x3a4b // TTY/TDD Phone
	pstItemFtpSite = 0x3a4c // Ftp Site
	pstItemGender = 0x3a4d // Gender
	pstItemManagersName = 0x3a4e // Manager's Name
	pstItemNickname = 0x3a4f // Nickname
	pstItemPersonalHomePage = 0x3a50 // Personal Home Page
	pstItemBusinessHomePage = 0x3a51 // Business Home Page
	pstItemCompanyMainPhone = 0x3a57 // Company Main Phone
	pstItemChildrenNames = 0x3a58 // childrens names
	pstItemHomeAddressCity = 0x3a59 // Home Address City
	pstItemHomeAddressCountry = 0x3a5a // Home Address Country
	pstItemHomeAddressPostalCode = 0x3a5b // Home Address Postal Code
	pstItemHomeAddressStateOrProvince = 0x3a5c // Home Address State or Province
	pstItemHomeAddressStreet = 0x3a5d // Home Address Street
	pstItemHomeAddressPostOfficeBox = 0x3a5e // Home Address Post Office Box
	pstItemOtherAddressCity = 0x3a5f // Other Address City
	pstItemOtherAddressCountry = 0x3a60 // Other Address Country
	pstItemOtherAddressPostalCode = 0x3a61 // Other Address Postal Code
	pstItemOtherAddressState = 0x3a62 // Other Address State
	pstItemOtherAddressStreet = 0x3a63 // Other Address Street
	pstItemOtherAddressPostOfficeBox = 0x3a64 // Other Address Post Office box
	pstItemInternetCodePage = 0x3fde // Internet code page
	pstItemMessageCodePage = 0x3ffd // Message code page
	pstItemEntryID = 0x65e3 // Entry ID
	pstItemAttachmentID2Value = 0x67f2 // Attachment ID2 value
	pstItemPasswordChecksum = 0x67ff // Password checksum
	pstItemSecureHTMLBody = 0x6f02 // Secure HTML Body
	pstItemSecureTextBody = 0x6f04 // Secure Text Body
	pstItemTopOfFoldersRecID = 0x7c07 // Top of folders RecID
	pstItemContactFullname = 0x8005 // Contact Fullname
	pstItemHomeAddress = 0x801a // Home Address
	pstItemBusinessAddress = 0x801b // Business Address
	pstItemOtherAddress = 0x801c // Other Address
	pstItemWorkAddressStreet = 0x8045 // Work Address Street
	pstItemWorkAddressCity = 0x8046 // Work Address City
	pstItemWorkAddressState = 0x8047 // Work Address State
	pstItemWorkAddressPostalCode = 0x8048 // Work Address Postal Code
	pstItemWorkAddressCountry = 0x8049 // Work Address Country
	pstItemWorkAddressPostOfficeBox = 0x804a // Work Address Post Office Box
	pstItemEmailAddress1Transport = 0x8082 // Email Address 1 Transport
	pstItemEmailAddress1Address = 0x8083 // Email Address 1 Address
	pstItemEmailAddress1Description = 0x8084 // Email Address 1 Description
	pstItemEmailAddress1Record = 0x8085 // Email Address 1 Record
	pstItemEmailAddress2Transport = 0x8092 // Email Address 2 Transport
	pstItemEmailAddress2Address = 0x8093 // Email Address 2 Address
	pstItemEmailAddress2Description = 0x8094 // Email Address 2 Description
	pstItemEmailAddress2Record = 0x8095 // Email Address 2 Record
	pstItemEmailAddress3Transport = 0x80a2 // Email Address 3 Transport
	pstItemEmailAddress3Address = 0x80a3 // Email Address 3 Address
	pstItemEmailAddress3Description = 0x80a4 // Email Address 3 Description
	pstItemEmailAddress3Record = 0x80a5 // Email Address 3 Record
	pstItemInternetFreeBusy = 0x80d8 // Internet Free/Busy
	pstItemAppointmentShowsAs = 0x8205 // Appointment shows as
	pstItemAppointmentLocation = 0x8208 // Appointment Location
	pstItemAppointmentStart = 0x820d // Appointment start
	pstItemAppointmentEnd = 0x820e // Appointment end
	pstItemLabelForAppointment = 0x8214 // Label for appointment
	pstItemAllDayAppointmentFlag = 0x8215 // All day appointment flag
	pstItemRecurrenceType = 0x8231 // Recurrence type
	pstItemRecurrenceDescription = 0x8232 // Recurrence description
	pstItemTimeZoneOfTimes = 0x8234 // TimeZone of times
	pstItemRecurrenceStartTime = 0x8235 // Recurrence Start Time
	pstItemRecurrenceEndTime = 0x8236 // Recurrence End Time
	pstItemReminderMinutesBeforeAppointmentStart = 0x8501 // Reminder minutes before appointment start
	pstItemReminderAlarm = 0x8503 // Reminder alarm
	pstItemCommonTimeStart = 0x8516 // Common Time Start
	pstItemCommonTimeEnd = 0x8517 // Common Time End
	pstItemPlayReminderSoundFilename = 0x851f // Play reminder sound filename
	pstItemFollowupString = 0x8530 // Followup String
	pstItemMileage = 0x8534 // Mileage
	pstItemBillingInformation = 0x8535 // Billing Information
	pstItemOutlookVersion = 0x8554 // Outlook Version
	pstItemAppointmentReminderTime = 0x8560 // Appointment Reminder Time
	pstItemJournalEntryType = 0x8700 // Journal Entry Type
	pstItemStartTimestamp = 0x8706 // Start Timestamp
	pstItemEndTimestamp = 0x8708 // End Timestamp
	pstItemJournalEntryType2 = 0x8712 // Journal Entry Type
)

func itemName(n uint16) string {
	switch n {
	case pstItemAlternateRecipientAllowed:
		return "Alternate recipient allowed"
	case pstItemExtendedAttributesTable:
		return "Extended Attributes Table"
	case pstItemImportanceLevel:
		return "Importance Level"
	case pstItemIPMContextMessageClass:
		return "IPM Context, message class"
	case pstItemGlobalDeliveryReportRequested:
		return "Global delivery report requested"
	case pstItemPriority:
		return "Priority"
	case pstItemReadReceipt:
		return "Read Receipt"
	case pstItemReassignmentProhibited:
		return "Reassignment Prohibited"
	case pstItemOriginalSensitivity:
		return "Original Sensitivity"
	case pstItemReportTime:
		return "Report time"
	case pstItemSensitivity:
		return "Sensitivity"
	case pstItemEmailSubject:
		return "Email Subject"
	case pstItemClientSubmitTimeDateSent:
		return "Client submit time / date sent"
	case pstItemOutlookAddressOfSender:
		return "Outlook Address of Sender"
	case pstItemOutlookStructureDescribingTheRecipient:
		return "Outlook structure describing the recipient"
	case pstItemNameOfTheOutlookRecipientStructure:
		return "Name of the Outlook recipient structure"
	case pstItemOutlookStructureDescribingTheSender:
		return "Outlook structure describing the sender"
	case pstItemNameOfTheOutlookSenderStructure:
		return "Name of the Outlook sender structure"
	case pstItemAnotherStructureDescribingTheRecipient:
		return "Another structure describing the recipient"
	case pstItemNameOfTheSecondRecipientStructure:
		return "Name of the second recipient structure"
	case pstItemReplyToOutlookStructure:
		return "Reply-To Outlook Structure"
	case pstItemNameOfTheReplyToStructure:
		return "Name of the Reply-To structure"
	case pstItemOutlookNameOfRecipient:
		return "Outlook Name of recipient"
	case pstItemSecondOutlookNameOfRecipient:
		return "Second Outlook name of recipient"
	case pstItemMyAddressInTOField:
		return "My address in TO field"
	case pstItemMyAddressInCCField:
		return "My address in CC field"
	case pstItemMessageAddressedToMe:
		return "Message addressed to me"
	case pstItemResponseRequested:
		return "Response requested"
	case pstItemSendersAddressAccessMethodSMTPEX:
		return "Sender's Address access method (SMTP, EX)"
	case pstItemSendersAddress:
		return "Sender's Address"
	case pstItemConversationTopicProcessedSubjectwithFwdReRemoved:
		return "Conversation topic, processed subject (with Fwd:, Re, ... removed)"
	case pstItemConversationIndex:
		return "Conversation index"
	case pstItemOriginalDisplayBCC:
		return "Original display BCC"
	case pstItemOriginalDisplayCC:
		return "Original display CC"
	case pstItemOriginalDisplayTO:
		return "Original display TO"
	case pstItemRecipientAddressAccessMethodSMTPEX:
		return "Recipient Address Access Method (SMTP, EX)"
	case pstItemRecipientsAddress:
		return "Recipient's Address"
	case pstItemSecondRecipientAccessMethodSMTPEX:
		return "Second Recipient Access Method (SMTP, EX)"
	case pstItemSecondRecipientAddress:
		return "Second Recipient Address"
	case pstItemEmailHeader:
		return "Email Header. This is the header that was attached to the email"
	case pstItemNDRReasonCode:
		return "NDR Reason code"
	case pstItemNDRDiagCode:
		return "NDR Diag code"
	case pstItemNonreceiptNotificationRequested:
		return "Non-receipt notification requested"
	case pstItemReplyRequested:
		return "Reply Requested"
	case pstItemSecondSenderStructure:
		return "Second sender structure"
	case pstItemNameOfSecondSenderStructure:
		return "Name of second sender structure"
	case pstItemSupplementaryInfo:
		return "Supplementary info"
	case pstItemSecondOutlookNameOfSender:
		return "Second outlook name of sender"
	case pstItemSecondSenderAccessMethodSMTPEX:
		return "Second sender access method (SMTP, EX)"
	case pstItemSecondSenderAddress:
		return "Second Sender Address"
	case pstItemNDRStatusCode:
		return "NDR status code"
	case pstItemDeleteAfterSubmit:
		return "Delete after submit"
	case pstItemBCCAddresses:
		return "BCC Addresses"
	case pstItemCCAddresses:
		return "CC Addresses"
	case pstItemSentToAddress:
		return "SentTo Address"
	case pstItemDate:
		return "Date."
	case pstItemFlagBits:
		return "Flag bits"
	case pstItemMessageSize:
		return "Message Size"
	case pstItemSentmailEntryID:
		return "Sentmail EntryID"
	case pstItemCompressedRTFInSync:
		return "Compressed RTF in Sync"
	case pstItemAttachmentSize:
		return "Attachment Size"
	case pstItembinaryRecordHeader:
		return "binary record header"
	case pstItemPlainTextEmailBody:
		return "Plain Text Email Body. Does not exist if the email doesn't have a plain text version"
	case pstItemReportText:
		return "Report Text"
	case pstItemRTFSyncBodyCRC:
		return "RTF Sync Body CRC"
	case pstItemRTFSyncBodyCharacterCount:
		return "RTF Sync Body character count"
	case pstItemRTFSyncBodyTag:
		return "RTF Sync body tag"
	case pstItemRTFCompressedBody:
		return "RTF Compressed body"
	case pstItemRTFWhitespacePrefixCount:
		return "RTF whitespace prefix count"
	case pstItemRTFWhitespaceTailingCount:
		return "RTF whitespace tailing count"
	case pstItemHTMLEmailBody:
		return "HTML Email Body. Does not exist if the email doesn't have an HTML version"
	case pstItemMessageID:
		return "Message ID"
	case pstItemInReplyToOrParentsMessageID:
		return "In-Reply-To or Parent's Message ID"
	case pstItemReturnPath:
		return "Return Path"
	case pstItemFolderName:
		return "Folder Name? I have seen this value used for the contacts record aswell"
	case pstItemAddressType:
		return "Address Type"
	case pstItemContactAddress:
		return "Contact Address"
	case pstItemComment:
		return "Comment"
	case pstItemDateItemCreation:
		return "Date item creation"
	case pstItemDateItemModification:
		return "Date item modification"
	case pstItembinaryRecordHeader2:
		return "binary record header"
	case pstItemValidFolderMask:
		return "Valid Folder Mask"
	case pstItemBinaryRecordContainsAReferenceToTopOfPersonalFolderItem:
		return "binary record contains a reference to \"Top of Personal Folder\" item"
	case pstItemBinaryRecordContainsAReferenceToDefaultOutboxItem:
		return "binary record contains a reference to default outbox item"
	case pstItemBinaryRecordContainsAReferenceToDeletedItemsItem:
		return "binary record contains a reference to \"Deleted Items\" item"
	case pstItemBinaryRecordContainsAReferenceToSentItemsFolderItem:
		return "binary record contains a reference to sent items folder item"
	case pstItemBinaryRecordContainsAReferenceToUserViewsFolderItem:
		return "binary record contains a reference to user views folder item"
	case pstItemBinaryRecordContainsAReferenceToCommonViewsFolderItem:
		return "binary record contains a reference to common views folder item"
	case pstItemBinaryRecordContainsAReferenceToSearchRootItem:
		return "binary record contains a reference to \"Search Root\" item"
	case pstItemNumberOfEmailsStoredInAFolder:
		return "the number of emails stored in a folder"
	case pstItemNumberOfUnreadEmailsInAFolder:
		return "the number of unread emails in a folder"
	case pstItemHasSubfolders:
		return "Has Subfolders"
	case pstItemFolderContentDescription:
		return "the folder content description"
	case pstItemAssociateContentCount:
		return "Associate Content count"
	case pstItemBinaryDataAttachment:
		return "Binary Data attachment"
	case pstItemAttachmentFilename:
		return "Attachment Filename"
	case pstItemAttachmentMethod:
		return "Attachement method"
	case pstItemAttachmentFilenameLong:
		return "Attachment Filename long"
	case pstItemAttachmentPosition:
		return "Attachment Position"
	case pstItemAttachmentMimeEncoding:
		return "Attachment mime encoding"
	case pstItemAttachmentMimeSequence:
		return "Attachment mime Sequence"
	case pstItemContactsAccountName:
		return "Contact's Account name"
	case pstItemContactAlternateRecipient:
		return "Contact Alternate Recipient"
	case pstItemCallbackTelephoneNumber:
		return "Callback telephone number"
	case pstItemMessageConversionProhibited:
		return "Message Conversion Prohibited"
	case pstItemContactsSuffix:
		return "Contacts Suffix"
	case pstItemContactsFirstName:
		return "Contacts First Name"
	case pstItemContactsGovernmentIDNumber:
		return "Contacts Government ID Number"
	case pstItemBusinessTelephoneNumber:
		return "Business Telephone Number"
	case pstItemHomeTelephoneNumber:
		return "Home Telephone Number"
	case pstItemContactsInitials:
		return "Contacts Initials"
	case pstItemKeyword:
		return "Keyword"
	case pstItemContactsLanguage:
		return "Contact's Language"
	case pstItemContactsLocation:
		return "Contact's Location"
	case pstItemMailPermission:
		return "Mail Permission"
	case pstItemMHSCommonName:
		return "MHS Common Name"
	case pstItemOrganizationalID:
		return "Organizational ID #"
	case pstItemContactsSurname:
		return "Contacts Surname"
	case pstItemoriginalEntryID:
		return "original entry id"
	case pstItemoriginalDisplayName:
		return "original display name"
	case pstItemoriginalSearchKey:
		return "original search key"
	case pstItemDefaultPostalAddress:
		return "Default Postal Address"
	case pstItemCompanyName:
		return "Company Name"
	case pstItemJobTitle:
		return "Job Title"
	case pstItemDepartmentName:
		return "Department Name"
	case pstItemOfficeLocation:
		return "Office Location"
	case pstItemPrimaryTelephone:
		return "Primary Telephone"
	case pstItemBusinessPhoneNumber2:
		return "Business Phone Number 2"
	case pstItemMobilePhoneNumber:
		return "Mobile Phone Number"
	case pstItemRadioPhoneNumber:
		return "Radio Phone Number"
	case pstItemCarPhoneNumber:
		return "Car Phone Number"
	case pstItemOtherPhoneNumber:
		return "Other Phone Number"
	case pstItemTransmittableDisplayName:
		return "Transmittable Display Name"
	case pstItemPagerPhoneNumber:
		return "Pager Phone Number"
	case pstItemuserCertificate:
		return "user certificate"
	case pstItemPrimaryFaxNumber:
		return "Primary Fax Number"
	case pstItemBusinessFaxNumber:
		return "Business Fax Number"
	case pstItemHomeFaxNumber:
		return "Home Fax Number"
	case pstItemBusinessAddressCountry:
		return "Business Address Country"
	case pstItemBusinessAddressCity:
		return "Business Address City"
	case pstItemBusinessAddressState:
		return "Business Address State"
	case pstItemBusinessAddressStreet:
		return "Business Address Street"
	case pstItemBusinessPostalCode:
		return "Business Postal Code"
	case pstItemBusinessPOBox:
		return "Business PO Box"
	case pstItemTelexNumber:
		return "Telex Number"
	case pstItemISDNNumber:
		return "ISDN Number"
	case pstItemAssistantPhoneNumber:
		return "Assistant Phone Number"
	case pstItemHomePhone2:
		return "Home Phone 2"
	case pstItemAssistantsName:
		return "Assistant's Name"
	case pstItemCanReceiveRichText:
		return "Can receive Rich Text"
	case pstItemWeddingAnniversary:
		return "Wedding Anniversary"
	case pstItemBirthday:
		return "Birthday"
	case pstItemHobbies:
		return "Hobbies"
	case pstItemMiddleName:
		return "Middle Name"
	case pstItemDisplayNamePrefixTitle:
		return "Display Name Prefix (Title)"
	case pstItemProfession:
		return "Profession"
	case pstItemPreferredByName:
		return "Preferred By Name"
	case pstItemSpousesName:
		return "Spouse's Name"
	case pstItemComputerNetworkName:
		return "Computer Network Name"
	case pstItemCustomerID:
		return "Customer ID"
	case pstItemTTYTDDPhone:
		return "TTY/TDD Phone"
	case pstItemFtpSite:
		return "Ftp Site"
	case pstItemGender:
		return "Gender"
	case pstItemManagersName:
		return "Manager's Name"
	case pstItemNickname:
		return "Nickname"
	case pstItemPersonalHomePage:
		return "Personal Home Page"
	case pstItemBusinessHomePage:
		return "Business Home Page"
	case pstItemCompanyMainPhone:
		return "Company Main Phone"
	case pstItemChildrenNames:
		return "childrens names"
	case pstItemHomeAddressCity:
		return "Home Address City"
	case pstItemHomeAddressCountry:
		return "Home Address Country"
	case pstItemHomeAddressPostalCode:
		return "Home Address Postal Code"
	case pstItemHomeAddressStateOrProvince:
		return "Home Address State or Province"
	case pstItemHomeAddressStreet:
		return "Home Address Street"
	case pstItemHomeAddressPostOfficeBox:
		return "Home Address Post Office Box"
	case pstItemOtherAddressCity:
		return "Other Address City"
	case pstItemOtherAddressCountry:
		return "Other Address Country"
	case pstItemOtherAddressPostalCode:
		return "Other Address Postal Code"
	case pstItemOtherAddressState:
		return "Other Address State"
	case pstItemOtherAddressStreet:
		return "Other Address Street"
	case pstItemOtherAddressPostOfficeBox:
		return "Other Address Post Office box"
	case pstItemInternetCodePage:
		return "Internet code page"
	case pstItemMessageCodePage:
		return "Message code page"
	case pstItemEntryID:
		return "Entry ID"
	case pstItemAttachmentID2Value:
		return "Attachment ID2 value"
	case pstItemPasswordChecksum:
		return "Password checksum"
	case pstItemSecureHTMLBody:
		return "Secure HTML Body"
	case pstItemSecureTextBody:
		return "Secure Text Body"
	case pstItemTopOfFoldersRecID:
		return "Top of folders RecID"
	case pstItemContactFullname:
		return "Contact Fullname"
	case pstItemHomeAddress:
		return "Home Address"
	case pstItemBusinessAddress:
		return "Business Address"
	case pstItemOtherAddress:
		return "Other Address"
	case pstItemWorkAddressStreet:
		return "Work Address Street"
	case pstItemWorkAddressCity:
		return "Work Address City"
	case pstItemWorkAddressState:
		return "Work Address State"
	case pstItemWorkAddressPostalCode:
		return "Work Address Postal Code"
	case pstItemWorkAddressCountry:
		return "Work Address Country"
	case pstItemWorkAddressPostOfficeBox:
		return "Work Address Post Office Box"
	case pstItemEmailAddress1Transport:
		return "Email Address 1 Transport"
	case pstItemEmailAddress1Address:
		return "Email Address 1 Address"
	case pstItemEmailAddress1Description:
		return "Email Address 1 Description"
	case pstItemEmailAddress1Record:
		return "Email Address 1 Record"
	case pstItemEmailAddress2Transport:
		return "Email Address 2 Transport"
	case pstItemEmailAddress2Address:
		return "Email Address 2 Address"
	case pstItemEmailAddress2Description:
		return "Email Address 2 Description"
	case pstItemEmailAddress2Record:
		return "Email Address 2 Record"
	case pstItemEmailAddress3Transport:
		return "Email Address 3 Transport"
	case pstItemEmailAddress3Address:
		return "Email Address 3 Address"
	case pstItemEmailAddress3Description:
		return "Email Address 3 Description"
	case pstItemEmailAddress3Record:
		return "Email Address 3 Record"
	case pstItemInternetFreeBusy:
		return "Internet Free/Busy"
	case pstItemAppointmentShowsAs:
		return "Appointment shows as"
	case pstItemAppointmentLocation:
		return "Appointment Location"
	case pstItemAppointmentStart:
		return "Appointment start"
	case pstItemAppointmentEnd:
		return "Appointment end"
	case pstItemLabelForAppointment:
		return "Label for appointment"
	case pstItemAllDayAppointmentFlag:
		return "All day appointment flag"
	case pstItemRecurrenceType:
		return "Recurrence type"
	case pstItemRecurrenceDescription:
		return "Recurrence description"
	case pstItemTimeZoneOfTimes:
		return "TimeZone of times"
	case pstItemRecurrenceStartTime:
		return "Recurrence Start Time"
	case pstItemRecurrenceEndTime:
		return "Recurrence End Time"
	case pstItemReminderMinutesBeforeAppointmentStart:
		return "Reminder minutes before appointment start"
	case pstItemReminderAlarm:
		return "Reminder alarm"
	case pstItemCommonTimeStart:
		return "Common Time Start"
	case pstItemCommonTimeEnd:
		return "Common Time End"
	case pstItemPlayReminderSoundFilename:
		return "Play reminder sound filename"
	case pstItemFollowupString:
		return "Followup String"
	case pstItemMileage:
		return "Mileage"
	case pstItemBillingInformation:
		return "Billing Information"
	case pstItemOutlookVersion:
		return "Outlook Version"
	case pstItemAppointmentReminderTime:
		return "Appointment Reminder Time"
	case pstItemJournalEntryType:
		return "Journal Entry Type"
	case pstItemStartTimestamp:
		return "Start Timestamp"
	case pstItemEndTimestamp:
		return "End Timestamp"
	case pstItemJournalEntryType2:
		return "Journal Entry Type"
	}
	return fmt.Sprintf("0x%04x", n)
}
